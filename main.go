package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Person struct {
	Name string `json:"name"`
}

func index(writer http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		var p Person
		json.NewDecoder(req.Body).Decode(&p)
		fmt.Fprintf(writer, "Hello World - %s", p.Name)
	case "GET":
		fmt.Fprint(writer, "Hello World with GET")
	default:
		writer.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func main() {
}
