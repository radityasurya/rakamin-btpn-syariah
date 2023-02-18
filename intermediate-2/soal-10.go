package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

func handleRequest() {
	http.HandleFunc("/articles", returnAllArticles)
	fmt.Println("starting web server at http://localhost:10000/")
	http.ListenAndServe(":10000", nil)
}

func main() {
	Articles = []Article{
		{Title: "Hello1", Desc: "Deskripsi artikel", Content: "Isi kontennya"},
		{Title: "Hello2", Desc: "Deskripsi artikel", Content: "Isi kontennya"},
	}

	handleRequest()
}
