package main

import (
	"fmt"
	"log"
	"net/http"

	"./service"
)

func main() {
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("assets"))))
	r := service.Router()

	fmt.Println("Starting server on the 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
