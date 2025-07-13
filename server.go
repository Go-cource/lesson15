package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	MainText := fmt.Sprintf("Это мой первый сайт!!\n Всем привет \nВремя: %v", time.Now().Format("02-01-06 15:04:05"))
	//http.ServeFile(w, r, "./static/page.html")
	tmpl, err := template.ParseFiles("./static/page.html")
	if err != nil {
		fmt.Println(err)
	}
	tmpl.Execute(w, MainText)
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
