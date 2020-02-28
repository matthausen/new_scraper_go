package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"net/http"
	"github.com/joho/godotenv"
)

func goDotEnvVariable(key string) string {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

type Response struct {
	Article []Article `json:"articles"`
}

type Article struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func main() {
	apiKey := goDotEnvVariable("API_KEY")
	var responseObject Response

	topics := []string{"china", "iraq", "bitcoin", "coronavirus", "ukraine", "iphone", "google", "syria", "virus", "global warming", "environment"}

	f, err := os.Create("train.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i < len(topics); i++ {
		resp, err := http.Get("https://newsapi.org/v2/everything?q=" + topics[i] + "&apiKey=" + apiKey)
		if err != nil {
			log.Fatalln(err)
		}

		defer resp.Body.Close()

		responseData, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			log.Fatalln(err)
		}
		json.Unmarshal(responseData, &responseObject)

		for i := 0; i < len(responseObject.Article); i++ {
			l, err := f.WriteString(responseObject.Article[i].Title + "\n" + responseObject.Article[i].Content)
			if err != nil {
				fmt.Println(err)
				f.Close()
				return
			}
			fmt.Println(l, "bytes written successfully")
		}
	}
	r := router.Router()
	fmt.Println("Starting server on the 8080...")

	log.Fatal(http.ListenAndServe(":8080", r))
}
