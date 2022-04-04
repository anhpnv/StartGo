package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Hello World</h1>")
}

func main() {
	fmt.Printf("Starting server at port 8081\n")
	http.HandleFunc("/", handler)
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
	// fmt.Println(syntax.Variable())
}
