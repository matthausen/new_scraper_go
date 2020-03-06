package service

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type Response struct {
	Article []Article `json:"articles"`
}

type Article struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func Router() *mux.Router {

	router := mux.NewRouter()
	router.HandleFunc("/", index)
	router.HandleFunc("/api/topic", getTopic).Methods("POST", "OPTIONS")
	return router
}

var tpl *template.Template
var endpoint = "http://newsapi.org/v2/top-headlines?sources=google-news&apiKey=API_KEY"
var apiKey = goDotEnvVariable("API_KEY")

func handleCors(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "*")
	(*w).Header().Set("Access-Control-Allow-Headers", "*")
	(*w).Header().Set("Content-Type", "*")
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func goDotEnvVariable(key string) string {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func fetchNews() {
	var responseObject Response

	resp, err := http.Get("http://newsapi.org/v2/top-headlines?sources=google-news&apiKey=" + apiKey)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	responseData, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}
	json.Unmarshal(responseData, &responseObject)
}

func index(res http.ResponseWriter, req *http.Request) {
	// fetchNews()
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
