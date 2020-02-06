package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"net/http"
	"encoding/json"
	"./config"
)

/* 
* Get headline by recency and source (e.g Google maps)
*	Get the top articles from targeted topics, at least 2 topics per categpory.
*	Write a .txt file with only the content from those topics
* 1 function to fetch the content and 1 function to write it
*/

// The final content
var c []string;

type Response struct {
	Article []Article `json:"articles"`
}

type Article struct {
	Description string `json:"description"`
	Content string `json:"content"`
}

func generateTrainData(c []string){
	f, err := os.Create("train_data.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i < len(c); i++ {
		l, err := f.WriteString(c[i])
		if err != nil {
			fmt.Println(err)
			f.Close()
			return
		}
		fmt.Println(l, "Bytes written successfully")
		err = f.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func main() {
	
	keyWords := []string{"bitcoin", "coronavirus", "ukraine"}
	for i := 0; i < len(keyWords); i++ {
		fmt.Println("My jkeyword: ", keyWords[i])
	}

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

	for i := 0; i < len(responseObject.Article); i++ {
		// fmt.Println(responseObject.Article[i].Content)
		c = append(c, responseObject.Article[i].Content)
		fmt.Println(c)
	}

	generateTrainData(c)

	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.ListenAndServe(":8080", nil)
}