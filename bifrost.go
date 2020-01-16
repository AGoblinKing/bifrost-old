package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Version of Bifrost
const Version = "0.0.5"

var port int64

func serveHTTP() {
	var portStr = strconv.FormatInt(port, 10)

	spa := spaHandler{
		staticPath: "docs",
		indexPath:  "index.html",
	}

	router := mux.NewRouter()

	router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		// an example API handler
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})

	router.PathPrefix("/").Handler(spa)

	srv := &http.Server{
		Handler:      router,
		Addr:         ":" + portStr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func main() {
	println("\u001b[32m", logo())

	flag.Int64Var(&port, "port", 5000, "port to bind to")
	flag.Parse()
	println("VERSION: ", Version)
	println("PORT: ", port)

	serveHTTP()
}
