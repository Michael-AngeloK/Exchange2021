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
	Country    string       `json:"name"`
	Currencies []Currency   `json:"currencies"`
	Border     []string     `json:"borders"`
	Exchange   ExchangeData `json:exchangedata`
}

type Currency struct {
	Code string `json:"code"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi, paste either of those on the back of the url:\n\n`/exchange/v1/exchangehistory/{country_name}`\n`/exchange/v1/exchangehistory/{country_name}/{begin_date-end_date}`\n`/exchange/v1/exchangehistory/{country_name}/{begin_date-end_date}`")

	fmt.Println("Endpoint Hit: homePage")
}

func exchangeHistory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	url := "https://restcountries.eu/rest/v2/name/" + params["country_name"]

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
	if responseObject[0].Currencies[0].Code == "EUR" {
		url1 += "&base=USD"
	}

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
	//fmt.Fprintf(w, string(url1))
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

func exchangeHistoryDates(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	url := "https://restcountries.eu/rest/v2/name/" + params["country_name"]

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
	if responseObject[0].Currencies[0].Code == "EUR" {
		url1 += "&base=USD"
	}

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

type ExchangeData struct {
	Rates Rate   `json:"rates"`
	Name  string `json:"name"`
	Base  string `json:"base"`
	Date  string `json:"date"`
}

type Rate struct {
	CAD float64 `json:"CAD,omitempty"`
	HKD float64 `json:"HKD,omitempty"`
	ISK float64 `json:"ISK,omitempty"`
	PHP float64 `json:"PHP,omitempty"`
	DKK float64 `json:"DKK,omitempty"`
	HUF float64 `json:"HUF,omitempty"`
	CZK float64 `json:"CZK,omitempty"`
	AUD float64 `json:"AUD,omitempty"`
	RON float64 `json:"RON,omitempty"`
	SEK float64 `json:"SEK,omitempty"`
	IDR float64 `json:"IDR,omitempty"`
	INR float64 `json:"INR,omitempty"`
	BRL float64 `json:"BRL,omitempty"`
	RUB float64 `json:"RUB,omitempty"`
	HRK float64 `json:"HRK,omitempty"`
	JPY float64 `json:"JPY,omitempty"`
	THB float64 `json:"THB,omitempty"`
	CHF float64 `json:"CHF,omitempty"`
	SGD float64 `json:"SGD,omitempty"`
	PLN float64 `json:"PLN,omitempty"`
	BGN float64 `json:"BGN,omitempty"`
	TRY float64 `json:"TRY,omitempty"`
	CNY float64 `json:"CNY,omitempty"`
	NOK float64 `json:"NOK,omitempty"`
	NZD float64 `json:"NZD,omitempty"`
	ZAR float64 `json:"ZAR,omitempty"`
	USD float64 `json:"USD,omitempty"`
	MXN float64 `json:"MXN,omitempty"`
	ILS float64 `json:"ILS,omitempty"`
	GBP float64 `json:"GBP,omitempty"`
	KRW float64 `json:"KRW,omitempty"`
	MYR float64 `json:"MYR,omitempty"`
}

type Final struct {
	Rate []Data `json:"rates"`
	Base string `json:"base"`
}

type Data struct {
	Name     string `json:"name"`
	Currency string `json:"currency"`
	Rate     Rate   `json:"rate"`
}

func exchangeBorder(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	url := "https://restcountries.eu/rest/v2/name/" + params["country_name"]

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

	Borders := ""

	for i := 0; i < len(responseObject[0].Border); i++ {
		Borders += responseObject[0].Border[i]
		if (i + 1) == len(responseObject[0].Border) {
		} else {
			Borders += ";"
		}
	}

	url1 := "https://restcountries.eu/rest/v2/alpha?codes=" + Borders

	//fmt.Println(url1)

	response1, err := http.Get(url1)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	defer response1.Body.Close()
	//bodyBytes, _ := ioutil.ReadAll(response1.Body)

	//bodyString := json.RawMessage(bodyBytes)

	url2 := "https://api.exchangeratesapi.io/latest?symbols=" + responseObject[0].Currencies[0].Code
	//url1 := "https://api.exchangeratesapi.io/history?start_at=2018-01-01&end_at=2018-09-01&symbols=ILS,JPY"

	response2, err := http.Get(url2)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData2, err := ioutil.ReadAll(response2.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject2 ExchangeData
	json.Unmarshal(responseData2, &responseObject2)

	responseObject[0].Exchange = responseObject2

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

	var data []Data

	data = append(data, Data{Name: responseObject[0].Country, Currency: responseObject[0].Currencies[0].Code, Rate: responseObject[0].Exchange.Rates})
	data = append(data, Data{Name: responseObject[0].Country, Currency: responseObject[0].Currencies[0].Code, Rate: responseObject[0].Exchange.Rates})

	final := Final{Rate: data, Base: responseObject[0].Exchange.Base}

	w.Header().Set("Content-Type", "application/json")

	//json.NewEncoder(w).Encode(responseData2)
	//json.NewEncoder(w).Encode(responseObject)
	//json.NewEncoder(w).Encode(data)
	json.NewEncoder(w).Encode(final)
	//json.NewEncoder(w).Encode(string(responseData1))
	//json.NewEncoder(w).Encode(bodyString)
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage).Methods("GET")
	myRouter.HandleFunc("/exchange/v1/exchangehistory/{country_name}", exchangeHistory).Methods("GET")
	myRouter.HandleFunc("/exchange/v1/exchangehistory/{country_name}/{begin_date-end_date}", exchangeHistoryDates).Methods("GET")
	myRouter.HandleFunc("/exchange/v1/exchangeborder/{country_name}", exchangeBorder).Methods("GET")
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
