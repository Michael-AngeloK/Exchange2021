package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"Desc"`
	Content string `json:"Content"`
}

type Articles []Article

func AllArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "Article 1 2", Desc: "Description 1", Content: " Content 1"},
		Article{Title: "Article 2", Desc: "Description 2", Content: " Content 2"},
		Article{Title: "Article 3", Desc: "Description 3", Content: " Content 3"},
	}
	fmt.Println("Endpoint Hint: All Articles Endpoint")

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(articles)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Prog2005, Here we go")

	fmt.Println("Endpoint Hit: homePage")
}

func exchangeHistory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	fmt.Fprintf(w, "exchangehistory")

	fmt.Println("Endpoint Hit: Exchange history")
}

func exchangeBorder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	fmt.Fprintf(w, "exchangeborder")

	fmt.Println("Endpoint Hit: Exchange Border")
}

func diag(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	fmt.Fprintf(w, "diag")

	fmt.Println("Endpoint Hit: Diagnostic")
}

func testPostArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Test POST endpoint worked")
}

func handelRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)
	/// We have two endpoints, for the main root, like localhost:4747, it runs homepage function and for localhost:4747/articles it executes AllArticles function
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", AllArticles).Methods("GET")
	myRouter.HandleFunc("/articles", testPostArticles).Methods("POST")
	myRouter.HandleFunc("/exchange/v1/exchangehistory", exchangeHistory).Methods("GET")
	myRouter.HandleFunc("/exchange/v1/exchangeborder", exchangeBorder).Methods("GET")
	myRouter.HandleFunc("/exchange/v1/diag", diag).Methods("GET")
	log.Fatal(http.ListenAndServe(getport(), myRouter))
}

func main() {
	handelRequests()
}

//// Get Port if it is set by environment, else use a defined one like "4747"
func getport() string {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return ":" + port
}
