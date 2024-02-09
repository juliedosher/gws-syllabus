package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	json.Unmarshal(jsonFile, &syllabi)

	http.HandleFunc("/hello-world", helloWorld)
	http.HandleFunc("/hello-world-json", helloWorldJson)
	http.HandleFunc("/syllabi", readAllSyllabi)
	http.HandleFunc("/read", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		fmt.Fprintf(w, "ID: %v", id)
	})

	http.HandleFunc("/hello-world-html", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	http.HandleFunc("/help", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "help.html")
	})

	http.HandleFunc("/", home)
	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my Go Web Server!")
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World - GWS")
}

func helloWorldJson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{“message” : ”Hello World - GWS”}`)
}

type Syllabus struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"lewisEmail"`
}

//go:embed data/all.json
var jsonFile []byte
var syllabi []Syllabus

func readAllSyllabi(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, string(jsonFile))
}

func readSyllabus(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, strconv.Itoa(len(syllabi)))
}
