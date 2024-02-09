package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// ID: sp24-cpsc-20000-002

func main() {
	json.Unmarshal(jsonFile, &syllabi)

	http.HandleFunc("/hello-world", helloWorld)
	http.HandleFunc("/hello-world-json", helloWorldJson)
	http.HandleFunc("/syllabi", readAllSyllabi)

	http.HandleFunc("/read", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		if syllabus, ok := getSyllabusFromId(id); ok {
			syllabusStr, _ := json.MarshalIndent(syllabus, "", "    ")
			fmt.Fprintf(w, string(syllabusStr))

		} else {
			fmt.Fprintf(w, "Syllabus with ID ‘%v’ not found", id)
		}
	})

	http.HandleFunc("/delete", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		if syllabus, ok := getSyllabusFromId(id); ok {
			fmt.Fprintf(w, "Delete request – stubbed")
			fmt.Fprintf(w, "\n\nSyllabus to be deleted:\n\n")
			syllabusStr, _ := json.MarshalIndent(syllabus, "", "    ")
			fmt.Fprintf(w, string(syllabusStr))

		} else {
			fmt.Fprintf(w, "Syllabus with ID ‘%v’ not found", id)
		}
	})

	http.HandleFunc("/update", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		if syllabus, ok := getSyllabusFromId(id); ok {
			fmt.Fprintf(w, "Update request – stubbed")
			fmt.Fprintf(w, "\n\nSyllabus to be updated:\n\n")
			syllabusStr, _ := json.MarshalIndent(syllabus, "", "    ")
			fmt.Fprintf(w, string(syllabusStr))

		} else {
			fmt.Fprintf(w, "Syllabus with ID ‘%v’ not found", id)
		}
	})

	http.HandleFunc("/create", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		if _, ok := getSyllabusFromId(id); ok {
			fmt.Fprintf(w, "Syllabus with ID ‘%v’ already exists", id)

		} else {
			fmt.Fprintf(w, "Create request – stubbed")
			fmt.Fprintf(w, "\nSyllabus to be created: %v", id)
		}
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

//go:embed data/syllabi.json
var jsonFile []byte
var syllabi []Syllabus

func readAllSyllabi(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, string(jsonFile))
}

func readSyllabus(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, strconv.Itoa(len(syllabi)))
}

func getSyllabusFromId(id string) (Syllabus, bool) {
	for _, syllabus := range syllabi {
		if syllabus.ID == id {
			return syllabus, true
		}
	}
	return Syllabus{}, false
}
