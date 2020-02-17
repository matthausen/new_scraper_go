package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"net/http"
)

/*
* The news aggregator needs to scrape information for the following entity types
MEDICAL
BIOLOGICAL
WEAPONS
TECHNOLOGY
*/

var API_KEY = "91d30a50507549e8a1134c0642992de4"

type Response struct {
	Article []Article `json:"articles"`
}

type Article struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func main() {
	
	keyWords := []string{"bitcoin", "coronavirus", "ukraine"}
	for i := 0; i < len(keyWords); i++ {
		fmt.Println("My jkeyword: ", keyWords[i])
	}

	var responseObject Response

	topics := []string{"china", "iraq", "bitcoin", "coronavirus", "ukraine", "iphone", "google", "syria", "virus", "global warming", "environment"}

	f, err := os.Create("train.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i < len(topics); i++ {
		resp, err := http.Get("https://newsapi.org/v2/everything?q=" + topics[i] + "&apiKey=" + API_KEY)
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
	/* http.Handle("/", http.FileServer(http.Dir("./static")))
	http.ListenAndServe(":8080", nil) */
}
