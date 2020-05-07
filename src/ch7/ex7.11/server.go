package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

var mu sync.Mutex // guards access to db

func (d database) list(w http.ResponseWriter, req *http.Request) {
	mu.Lock()
	for item, price := range d {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
	mu.Unlock()
}

func (d database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	mu.Lock()
	price, ok := d[item]
	mu.Unlock()
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

func (d database) update(w http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	item := query.Get("item")
	price, err := strconv.ParseFloat(query.Get("price"), 32)
	if err != nil || price <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "invalid price: %q\n", query.Get("price"))
		return
	}
	mu.Lock()
	defer mu.Unlock()
	if _, ok := d[item]; !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	d[item] = dollars(price)
}

func main() {
	db := database{"shoes": 100, "pants": 20}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/update", db.update)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
