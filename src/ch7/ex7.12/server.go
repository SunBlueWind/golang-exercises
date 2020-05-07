package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (d database) list(w http.ResponseWriter, req *http.Request) {
	var template = template.Must(template.New("list").Parse(`
	<table>
	<tr>
		<th>Item</th>
		<th>Price</th>
	</tr>
	{{range $item, $price := .}}
	<tr>
		<td>{{$item}}</td>
		<td>{{$price}}</td>
	</tr>
	{{end}}
	</table>
	`))
	template.Execute(w, d)
}

func (d database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := d[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

func main() {
	db := database{"shoes": 100, "pants": 20}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
