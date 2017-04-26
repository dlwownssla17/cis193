package server

import (
	"net/http"
	"html/template"
	"coderive/src/indexer"
)

// Result represents the search result following a query.
type Result struct {
	Query string
	Elapsed int64
	Matches []*indexer.Match
}

/* * */

func homeHandler(w http.ResponseWriter, _ *http.Request) {
	renderTemplate(w, "home", nil)
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "search", nil)
}

func helpHandler(w http.ResponseWriter, _ *http.Request) {
	renderTemplate(w, "help", nil)
}

/* * */

var templates = template.Must(template.ParseFiles("../view/home.html", "../view/search.html", "../view/help.html"))

func renderTemplate(w http.ResponseWriter, template string, res *Result) {
	err := templates.ExecuteTemplate(w, template + ".html", res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

/* * */

// RunServer runs the server.
func RunServer() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/search/", searchHandler)
	http.HandleFunc("/help/", helpHandler)

	http.ListenAndServe(":8080", nil)
}