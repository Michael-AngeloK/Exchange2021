package RESTstudent

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// replyWithAllStudents prepares a response with all students from the student storage
func replyWithAllStudents(w io.Writer, db StudentsStorage) {
	if db.Count() == 0 {
		err := json.NewEncoder(w).Encode([]Student{})
		if err != nil {
			// this should never happen
			fmt.Println("ERROR encoding JSON for an empty array", err)
		}
	} else {
		a := make([]Student, 0, db.Count())
		a = append(a, db.GetAll()...)
		err := json.NewEncoder(w).Encode(a)
		if err != nil {
			fmt.Println("ERROR encoding JSON", err)
		}
	}
}

// replyWithStudent prepares a response with a single student from the student storage
func replyWithStudent(w http.ResponseWriter, db StudentsStorage, id string) {
	// make sure that i is valid
	s, ok := db.Get(id)
	if !ok {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	// handle /student/<id>
	err := json.NewEncoder(w).Encode(s)
	if err != nil {
		fmt.Println("ERROR encoding JSON", err)
	}
}

// HandlerStudent main handler for route related to `/student` requests
// Note: we using here a higher-order function with closure, to propagate a reference
// to the DB down the processing pipeline.
func HandlerStudent(db StudentsStorage) func (http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			handleStudentPost(w, r, db)
		case http.MethodGet:
			handleStudentGet(w, r, db)
		}
	}
}

// handleStudentPost utility function, package level, for handling POST request
func handleStudentPost(w http.ResponseWriter, r *http.Request, db StudentsStorage) {
	var s Student
	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Println("Decoding: " + err.Error())
		return
	}
	// check if the student is new
	_, ok := db.Get(s.StudentID)
	if ok {
		// TODO find a better Error Code (HTTP Status)
		http.Error(w, "Student already exists. Use PUT to modify.", http.StatusBadRequest)
		fmt.Println("Student already exists.")
		return
	}
	// new student
	fmt.Println("Adding to db ...")
	err = db.Add(s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println(http.StatusInternalServerError)
		return
	}
	_, err = fmt.Fprint(w, http.StatusOK) // 200 by default
	if err != nil {
		// TODO need to handle the error!
		// log it? print it? panic?
	}
}

// handleStudentGet utility function, package level, to handle GET request to student route
func handleStudentGet(w http.ResponseWriter, r *http.Request, db StudentsStorage) {
		http.Header.Add(w.Header(), "content-type", "application/json")
		// alternative way:
		// w.Header().Add("content-type", "application/json")
		parts := strings.Split(r.URL.Path, "/")
		// error handling
		if len(parts) != 3 || parts[1] != "student" {
			http.Error(w, "Malformed URL", http.StatusBadRequest)
			return
		}
		// handle the request /student/  which will return ALL students as array of JSON objects
		if parts[2] == "" {
			replyWithAllStudents(w, db)
		} else {
			replyWithStudent(w, db, parts[2])
		}
}

