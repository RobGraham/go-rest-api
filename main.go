package main

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
	"runtime"

	"github.com/RobGraham/go-rest-api/model"
	"github.com/gorilla/mux"
)

func main() {
	runtime.GOMAXPROCS(8)
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/todos", todoIndex)
	router.HandleFunc("/todos/{todoId}", todoShow)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func todoIndex(w http.ResponseWriter, r *http.Request) {
	todos := model.Todos{
		model.Todo{Name: "Hire presentation"},
		model.Todo{Name: "Host meetup"},
	}

	err := json.NewEncoder(w).Encode(todos)

	if err != nil {
		//handle err
	}
}

func todoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoID := vars["todoId"]

	fmt.Fprintln(w, "Showing todo:", todoID)
}
