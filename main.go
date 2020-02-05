package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"encoding/json"
	"./config"
)

type Response struct {
	Article []Article `json:"articles"`
}

type Article struct {
	Content string `json:"content"`
}


func main() {

	resp, err := http.Get(config.TopHeadlines)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	responseData, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}	

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	fmt.Println(responseObject.Article)

	for i := 0; i < len(responseObject.Article); i++ {
		fmt.Println(responseObject.Article[i].Content)
	}

	// log.Println(string(body))

	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.ListenAndServe(":8080", nil)
}