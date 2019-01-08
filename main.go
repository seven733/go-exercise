package main

import (
	"net/http"
	"time"

	"go-exercise/src/db"

	"go-exercise/src/route"
)

func init() {
	db.Connect()
}

func main() {
	server := &http.Server{
		Addr:         ":2333",
		Handler:      route.InitRoute(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server.ListenAndServe()
}
