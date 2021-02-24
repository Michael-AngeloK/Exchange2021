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

// A Response struct to restcountries
type Response struct {
	Country    string     `json:"name"`
	Currencies []currency `json:"currencies"`
	Border     []string   `json:"borders"`
}

type currency struct {
	Code string `json:"code"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Prog2005, Here we go")

	fmt.Println("Endpoint Hit: homePage")
}

func exchangeHistory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	url := "https://restcountries.eu/rest/v2/name/" + params["name"]

	response, err := http.Get(url)
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

	url1 := "https://api.exchangeratesapi.io/latest?symbols=" + responseObject[0].Currencies[0].Code
	//url1 := "https://api.exchangeratesapi.io/history?start_at=2018-01-01&end_at=2018-09-01&symbols=ILS,JPY"

	response1, err := http.Get(url1)
	if err != nil {
		log.Fatalln(err)
	}

	defer response1.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(response1.Body)

	bodyString := json.RawMessage(bodyBytes)
	//fmt.Println("API Response as String:\n" + bodyString)

	//fmt.Println(responseData1)
	//fmt.Println(len(responseObject.Borders))
	//fmt.Println(responseObject1)
	//fmt.Fprintf(w, string(responseData1))
	//delim := ""
	//fmt.Fprintf(w, string((strings.Trim(strings.Join(strings.Fields(fmt.Sprint(responseObject[0].Currencies)), delim), "[]"))))
	//fmt.Fprintf(w, string(responseData1))
	//for i := 0; i < len(responseObject[0].Border); i++ {
	//	fmt.Fprintln(w, responseObject[0].Border[i])
	//}

	w.Header().Set("Content-Type", "application/json")

	//json.NewEncoder(w).Encode(responseObject)
	//json.NewEncoder(w).Encode(responseObject1)
	//json.NewEncoder(w).Encode(string(responseData1))
	json.NewEncoder(w).Encode(bodyString)

	myString := "2020-12-01-2021-01-31"

	// Step 1: Convert it to a rune
	a := []rune(myString)
	//"2020-12-01-2021-01-31"
	// Step 2: Grab the num of chars you need
	myShortString := string(a[:10])
	myShortString1 := string(a[11:])

	fmt.Println(myShortString)
	fmt.Println(myShortString1)
}

func exchangeHistoryDates(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	url := "https://restcountries.eu/rest/v2/name/" + params["name"]

	response, err := http.Get(url)
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

	myString := params["begin_date-end_date"]

	// Step 1: Convert it to a rune
	a := []rune(myString)
	//"2020-12-01-2021-01-31"
	// Step 2: Grab the num of chars you need
	BeginDate := string(a[:10])
	EndDate := string(a[11:])

	url1 := "https://api.exchangeratesapi.io/history?start_at=" + BeginDate + "&end_at=" + EndDate + "&symbols=" + responseObject[0].Currencies[0].Code

	response1, err := http.Get(url1)
	if err != nil {
		log.Fatalln(err)
	}

	defer response1.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(response1.Body)

	bodyString := json.RawMessage(bodyBytes)
	//fmt.Println("API Response as String:\n" + bodyString)

	//fmt.Println(responseData1)
	//fmt.Println(len(responseObject.Borders))
	//fmt.Println(responseObject1)
	//fmt.Fprintf(w, string(responseData1))
	//delim := ""
	//fmt.Fprintf(w, string((strings.Trim(strings.Join(strings.Fields(fmt.Sprint(responseObject[0].Currencies)), delim), "[]"))))
	//fmt.Fprintf(w, string(responseData1))
	//for i := 0; i < len(responseObject[0].Border); i++ {
	//	fmt.Fprintln(w, responseObject[0].Border[i])
	//}

	w.Header().Set("Content-Type", "application/json")

	//json.NewEncoder(w).Encode(responseObject)
	//json.NewEncoder(w).Encode(responseObject1)
	//json.NewEncoder(w).Encode(string(responseData1))
	json.NewEncoder(w).Encode(bodyString)
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage).Methods("GET")
	myRouter.HandleFunc("/exchange/v1/exchangehistory/{name}", exchangeHistory).Methods("GET")
	myRouter.HandleFunc("/exchange/v1/exchangehistory/{name}/{begin_date-end_date}", exchangeHistoryDates).Methods("GET")
	log.Fatal(http.ListenAndServe(getport(), myRouter))
}

func main() {
	handleRequests()
}

func getport() string {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return ":" + port
}
