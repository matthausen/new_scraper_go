package handlers

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

var tpl *template.Template

func handleCors(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "*")
	(*w).Header().Set("Access-Control-Allow-Headers", "*")
	(*w).Header().Set("Content-Type", "*")
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func index(res http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(res, "index.gohtml", nil)
	if err != nil {
		log.Fatalln("template didn't execute: ", err)
	}
}

// GetTopic : get the user input on news topic
func getTopic(w http.ResponseWriter, r *http.Request) {
	handleCors(&w, r)

	if (*r).Method == "OPTIONS" {
		return
	}
	switch r.Method {
	case "GET":
		fmt.Fprintf(w, "Request method not supported.")
	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		fmt.Fprintf(w, "r.PostForm = %v\n", r.PostForm)
		topic := r.FormValue("topic")
		fmt.Fprintf(w, "Topic = %s\n", topic)

	default:
		fmt.Fprintf(w, "Request method not supported.")
	}
}
