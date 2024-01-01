package main

import (
	"fmt"
	"log"
	"net/http"
)

type dollars float32

type database map[string]dollars

func (db database) list(w http.ResponseWriter, rqe *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: $%.2f\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "$%.2f\n", price)
}

func main() {
	db := database{"shoes": 50, "socks": 5}
	mux := http.NewServeMux()
	mux.HandleFunc("/list", db.list)
	mux.HandleFunc("/price", db.price)
	// net/http package provides a global ServeMux instance called
	// Default ServerMux and package level functions called http.Handle
	// and http.HandleFunc. To use DefaultServeMux as the server's main
	// handler, we needn't pass it to ListenAndServe; nil will do like
	// http.ListenAndServe("localhost:8080", nil)
	// Ref: page 195 of "The Go Programming Language" book
	log.Fatal(http.ListenAndServe("localhost:8080", mux))
}
