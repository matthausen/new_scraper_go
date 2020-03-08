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
	Image   string `json:"urlToImage"`
	URL     string `json:"url"`
	Author  string `json:"author"`
}

type New struct {
	Title  string
	Text   string
	Image  string
	URL    string
	Author string
}

var tpl *template.Template
var endpoint = "http://newsapi.org/v2/top-headlines?sources=google-news&apiKey=API_KEY"
var apiKey = goDotEnvVariable("API_KEY")

func Router() *mux.Router {

	router := mux.NewRouter()
	router.HandleFunc("/", index).Methods("GET", "POST", "OPTIONS")
	router.HandleFunc("/topic", byTopic).Methods("GET", "POST", "OPTIONS")
	return router
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func handleCors(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "*")
	(*w).Header().Set("Access-Control-Allow-Headers", "*")
	(*w).Header().Set("Content-Type", "*")
}

func goDotEnvVariable(key string) string {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func fetchNews() []New {
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

	news := make([]New, 0)

	for i := 0; i < len(responseObject.Article); i++ {
		new := New{
			responseObject.Article[i].Title,
			responseObject.Article[i].Content,
			responseObject.Article[i].Image,
			responseObject.Article[i].URL,
			responseObject.Article[i].Author,
		}
		news = append(news, new)
		if err != nil {
			fmt.Println(err)
		}
	}
	return news
}

func fetchByTopic(t string) []New {
	var responseObject Response

	resp, err := http.Get("http://newsapi.org/v2/everything?q=" + t + "&apiKey=" + apiKey)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	responseData, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}
	json.Unmarshal(responseData, &responseObject)

	news := make([]New, 0)

	for i := 0; i < len(responseObject.Article); i++ {
		new := New{
			responseObject.Article[i].Title,
			responseObject.Article[i].Content,
			responseObject.Article[i].Image,
			responseObject.Article[i].URL,
			responseObject.Article[i].Author,
		}
		news = append(news, new)
		if err != nil {
			fmt.Println(err)
		}
	}
	return news
}

func byTopic(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Println("GET request in action")
	case "POST":
		if err := r.ParseForm(); err != nil {
			r.ParseForm()
		}
		topic := r.FormValue("topic")
		news := fetchByTopic(topic)
		err := tpl.ExecuteTemplate(w, "bytopic.gohtml", news)
		if err != nil {
			log.Fatalln("template didn't execute: ", err)
		}
	default:
		fmt.Fprintf(w, "Request method not supported.")
	}
}

func index(w http.ResponseWriter, r *http.Request) {

	news := fetchNews()
	handleCors(&w, r)

	if (*r).Method == "OPTIONS" {
		return
	}

	err := tpl.ExecuteTemplate(w, "index.gohtml", news)
	if err != nil {
		log.Fatalln("template didn't execute: ", err)
	}
}
