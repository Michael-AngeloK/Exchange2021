package main

import (
	"RESTstudent"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)


func handlerHello(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 4 {
		status := http.StatusBadRequest
		http.Error(w, "Expecting format .../firstname/lastname", status)
		return
	}
	name := parts[2]
	_, err := fmt.Fprintln(w, parts)
	if err != nil {
		// TODO must handle the error!
	}
	_, err = fmt.Fprintf(w, "Hello %s %s!\n", name, parts[3])
	if err != nil {
		// TODO must handle the error!
	}
}

// -----------------

func main() {
	// DB init
	db := RESTstudent.InitStudentsStorage()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/hello/", handlerHello)
	http.HandleFunc("/student/", RESTstudent.HandlerStudent(db)) // ensure to type complete URL when requesting
	fmt.Println("Listening on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
