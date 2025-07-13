package main

import (
	"fmt"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Index Page")
}
func AboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "About Page")
}
func ContactsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "About Page")
}

func main() {

	fmt.Println("Sever is listening...")
	http.ListenAndServe("127.0.0.1:8080", nil)
}
