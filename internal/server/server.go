package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/DanielKirkwood/youchooseserver/config"
	"github.com/DanielKirkwood/youchooseserver/internal/middleware"
	db "github.com/DanielKirkwood/youchooseserver/third_party/database"
)

// A Server holds all the modules required for the rest API.
// It can be extended when new modules are required such as
// a DB object for storing dbsn.
type Server struct {
	Version    string
	cfg        *config.Config
	DB         *sqlx.DB
	router     *chi.Mux
	httpServer *http.Server
}

type Options func(opts *Server) error

// New returns a new Server with the given optional
// options
func New(opts ...Options) *Server {
	s := defaultServer()

	for _, opt := range opts {
		err := opt(s)
		if err != nil {
			log.Fatalln(err)
		}
	}
	return s
}

// defaultServer returns a new basic Server struct.
// Can be easily updated when more custom configs are added.
func defaultServer() *Server {
	return &Server{
		cfg:    config.New(),
		router: chi.NewRouter(),
	}
}

// Init initiales the server for which it is called for.
// Each initialisation should be called from here, but the
// implementation should be carried out in another function
// to keep this init function simple and clear.
func (s *Server) Init(version string) {
	s.Version = version
	s.newRouter()
	s.newDatabase()
	s.setGlobalMiddleware()
}

// newRouter creates a new chi router on the servers
// router object.
func (s *Server) newRouter() {
	s.router = chi.NewRouter()
}

func (s *Server) newDatabase() {
	if s.cfg.Database.Driver == "" {
		log.Fatal("please fill in database credentials in .env file or set in environment variable")
	}
	s.DB = db.NewSqlx(s.cfg.Database)
	s.DB.SetMaxOpenConns(s.cfg.Database.MaxConnectionPool)
	s.DB.SetMaxIdleConns(s.cfg.Database.MaxIdleConnections)
	s.DB.SetConnMaxLifetime(s.cfg.Database.ConnectionsMaxLifeTime)
}

// setGlobalMiddleware enables our custom middleware on
// the server's router.
func (s *Server) setGlobalMiddleware() {
	s.router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(`{"message": "endpoint not found"}`))
	})
	s.router.Use(middleware.Json)
	s.router.Use(middleware.Audit)
	s.router.Use(middleware.CORS)
	if s.cfg.Api.RequestLog {
		s.router.Use(chiMiddleware.Logger)
	}
}

// Run runs the server.
func (s *Server) Run() {
	s.httpServer = &http.Server{
		Addr:              s.cfg.Api.Host + ":" + s.cfg.Api.Port,
		Handler:           s.router,
		ReadHeaderTimeout: s.cfg.Api.ReadHeaderTimeout,
	}

	fmt.Println(` __     __            _____ _                             _____
 \ \   / /           / ____| |                           / ____|
  \ \_/ ___  _   _  | |    | |__   ___   ___  ___  ___  | (___   ___ _ ____   _____ _ __
   \   / _ \| | | | | |    | '_ \ / _ \ / _ \/ __|/ _ \  \___ \ / _ | '__\ \ / / _ | '__|
    | | (_) | |_| | | |____| | | | (_) | (_) \__ |  __/  ____) |  __| |   \ V |  __| |
    |_|\___/ \__,_|  \_____|_| |_|\___/ \___/|___/\___| |_____/ \___|_|    \_/ \___|_|`)

	go func() {
		start(s)
	}()

	_ = gracefulShutdown(context.Background(), s)
}

// Config returns the server cfg.
func (s *Server) Config() *config.Config {
	return s.cfg
}

// start serves the given server.
func start(s *Server) {
	log.Printf("Serving at %s:%s\n", s.cfg.Api.Host, s.cfg.Api.Port)
	err := s.httpServer.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}

// gracefulShutdown shutsdown the server when it is killed
// by either CTRL-C or other means.
func gracefulShutdown(ctx context.Context, s *Server) error {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	log.Println("Shutting down...")

	ctx, shutdown := context.WithTimeout(ctx, s.Config().Api.GracefulTimeout*time.Second)
	defer shutdown()

	err := s.httpServer.Shutdown(ctx)
	if err != nil {
		log.Println(err)
	}

	return nil
}
