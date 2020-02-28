package handlers

import (
	"fmt"
	"net/http"
)

func handleCors(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "*")
	(*w).Header().Set("Access-Control-Allow-Headers", "*")
	(*w).Header().Set("Content-Type", "*")
}

func GetTopic(w http.ResponseWriter, r *http.Request) {
	handleCors(&w, r)

	if (*r).Method == "OPTIONS" {
		return
	}
	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "form.html")
	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		fmt.Fprintf(w, "Post from website! r.PostForm = %v\n", r.PostForm)
		topic := r.FormValue("topic")
		fmt.Fprintf(w, "Topic = %s\n", topic)

	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}