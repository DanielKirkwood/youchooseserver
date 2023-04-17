package database

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"

	"github.com/DanielKirkwood/youchooseserver/config"
)

// NewSqlx returns a new sqlx database using
// the given database config
func NewSqlx(cfg config.Database) *sqlx.DB {
	var dsn string

	if cfg.Driver == "postgres" || cfg.Driver == "pgx" {
		dsn = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			cfg.Host, cfg.Port, cfg.User, cfg.Pass, cfg.Name)
	} else {
		log.Fatal("Must choose postgres as dsn")
	}

	db, err := sqlx.Open(cfg.Driver, dsn)
	if err != nil {
		log.Fatalln(err)
	}

	alive(db.DB)

	return db
}

// alive connects to the database.
// Replaces Ping which is un-reliable as the connection
// is cached.
func alive(db *sql.DB) {
	log.Println("connecting to database...")
	for {
		_, err := db.Exec("SELECT true")
		if err == nil {
			log.Println("database connected")
			return
		}

		base, capacity := time.Second, time.Minute
		for backoff := base; err != nil; backoff <<= 1 {
			if backoff > capacity {
				backoff = capacity
			}

			// A pseudo-random number generator here is fine. No need to be
			// cryptographically secure. Ignore with the following comment:
			/* #nosec */
			jitter := rand.Int63n(int64(backoff * 3))
			sleep := base + time.Duration(jitter)
			time.Sleep(sleep)
			_, err := db.Exec("SELECT true")
			if err == nil {
				log.Println("database connected")
				return
			}
		}
	}
}
