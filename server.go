package main

import (
	"fmt"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/page.html")
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/about.html")
}

func ContactsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Contacts Page")
}

func main() {

	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/about", AboutHandler)
	http.HandleFunc("/contacts", ContactsHandler)

	fmt.Println("Server is listening...")
	http.ListenAndServe("127.0.0.1:8080", nil)
}
