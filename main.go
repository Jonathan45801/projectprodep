package main

import
(
	"fmt"
	"log"
	"net/http"
	"encoding/json"
)

type article struct{
	Title string `json:"Title"`
    Desc string `json:"desc"`
    Content string `json:"content"`
}
var articles []article
func returnAllArticles(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: returnAllArticles")
    json.NewEncoder(w).Encode(articles)
}
func homepage(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w,"Home Page Endpoint Hit")
}
func handRequest() {
	http.HandleFunc("/",homepage)
	http.HandleFunc("/articles", returnAllArticles)
	log.Fatal(http.ListenAndServe(":1000",nil))
}
func main()  {
	articles = []article{
        article{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
        article{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
    }
	handRequest();
}
