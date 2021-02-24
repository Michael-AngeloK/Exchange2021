package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// A Response struct to map the Entire Response
type Response struct {
	Country string   `json:"name"`
	Border  []string `json:"borders"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Prog2005, Here we go")

	fmt.Println("Endpoint Hit: homePage")
}

func pokemonGo(w http.ResponseWriter, r *http.Request) {
	response, err := http.Get("https://restcountries.eu/rest/v2/name/norway")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject []Response
	json.Unmarshal(responseData, &responseObject)

	//fmt.Println(responseData)
	//fmt.Println(len(responseObject.Borders))
	fmt.Println(responseObject)
	fmt.Fprintf(w, string(responseData))
	for i := 0; i < len(responseObject[0].Border); i++ {
		fmt.Println(responseObject[0].Border[i])
	}
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage).Methods("GET")
	myRouter.HandleFunc("/bob", pokemonGo).Methods("GET")
	log.Fatal(http.ListenAndServe(getport(), myRouter))
}

func main() {
	handleRequests()
}

func getport() string {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	return ":" + port
}
