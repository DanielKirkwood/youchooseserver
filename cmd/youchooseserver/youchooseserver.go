package main

import (
	"log"
)

// Version is injected using ldflags during build time
var Version = "v0.1.0"

// @title youchooseserver
// @version 0.1.0
// @description Go server for You Choose mobile application.
// @contact.name DanielKirkwood
// @contact.url https://github.com/DanielKirkwood/youchooseserver
// @contact.email danielkirkwood1973@gmail.com
// @host localhost:3080
// @BasePath /
func main() {
	log.Printf("Starting API version: %s\n", Version)
}
