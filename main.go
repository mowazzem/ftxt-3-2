package main

import (
	"ftxt-3-2/flag"
	"ftxt-3-2/login"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	loginHandler := login.NewLoginHandler()
	flagHandler := flag.NewFlagHandler()
	r := mux.NewRouter()
	r.HandleFunc("/login", loginHandler.Login).Methods("PUT")
	r.HandleFunc("/flag", flagHandler.PutFlag).Methods("PUT")
	r.HandleFunc("/flag", flagHandler.GetFlag).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	srv := &http.Server{
		Handler: r,
		Addr:    ":" + port,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
