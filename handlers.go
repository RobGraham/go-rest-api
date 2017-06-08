package main

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"
	"strconv"

	"github.com/RobGraham/go-rest-api/model"
	"github.com/gorilla/mux"
)

var todos = model.Todos{
	model.Todo{Name: "Hire presentation"},
	model.Todo{Name: "Host meetup"},
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {

	if err := json.NewEncoder(w).Encode(todos); err != nil {
		fmt.Println(err)
	}
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoID, err := strconv.Atoi(vars["todoId"])

	if err != nil {
		defer fmt.Fprintln(w, "Todo Hanlder Error", err)
	}

	if todoID > len(todos) || todoID < 1 {
		fmt.Fprintln(w, "No Todo found")
	} else {
		if _, err = fmt.Fprintln(w, "Showing todo:", todos[todoID-1]); err != nil {
			panic(err)
		}
	}
}
