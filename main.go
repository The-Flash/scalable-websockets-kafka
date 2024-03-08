package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func roomHandler(w http.ResponseWriter, r *http.Request) {
	templ := template.Must(template.ParseFiles("room.html"))
	templ.Execute(w, nil)
}

func main() {
	port := "8080"
	mux := http.NewServeMux()

	mux.HandleFunc("GET /room/{id}", roomHandler)

	fmt.Printf("Server is running on port %s\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), mux); err != nil {
		fmt.Println(err)
	}
}
