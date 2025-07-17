package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

type NewData struct {
	Article string
	Message string
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	MainText := fmt.Sprintf("Это мой первый сайт!!\n Всем привет \nВремя: %v", time.Now().Format("02-01-06 15:04:05"))
	data := NewData{
		Article: MainText,
		Message: "This is new message",
	}
	//http.ServeFile(w, r, "./static/page.html")
	tmpl, err := template.ParseFiles("./static/page.html")
	if err != nil {
		fmt.Println(err)
	}
	tmpl.Execute(w, data)
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/about.html")
}

func ContactsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Contacts Page")
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", IndexHandler)
	mux.HandleFunc("/about", AboutHandler)
	mux.HandleFunc("/about/contacts/", ContactsHandler)

	fmt.Println("Server is listening...")
	http.ListenAndServe("127.0.0.1:8080", mux)
}
