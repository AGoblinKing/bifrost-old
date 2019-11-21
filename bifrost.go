package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Version of Bifrost
const Version = "0.0.3"

var port int64

func serveHTTP() {
	var portStr = strconv.FormatInt(port, 10)

	spa := spaHandler{
		staticPath: "docs",
		indexPath:  "index.html",
	}

	r := mux.NewRouter()

	r.Headers("Cache-Control", "no-cache").
		Handler(spa)

	srv := &http.Server{
		Handler:      r,
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
