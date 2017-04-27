package server

import (
	"net/http"
	"html/template"
	"coderive/src/indexer"
	"runtime"
	"path"
	"time"
)

// Result represents the search result following a query.
type Result struct {
	Query string
	Matches []*indexer.Match
	StartTime time.Time
	EndTime time.Time
}

/* * */

func redirectHome(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "", http.StatusFound)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Path) > 1 {
		redirectHome(w, r)
		return
	}

	if r.Method == "GET" {
		renderTemplate(w, "home", nil)
	} else if r.Method == "POST" {
		redirectHome(w, r)
	} else {
		redirectHome(w, r)
	}
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	if r.Form.Get("q") == "" {
		redirectHome(w, r)
		return
	}

	res := &Result{
		Query: r.Form.Get("q"),
		StartTime: time.Now(),
	}



	if r.Method == "GET" {
		renderTemplate(w, "search", res)
	} else if r.Method == "POST" {
		redirectHome(w, r)
	} else {
		redirectHome(w, r)
	}
}

func helpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		renderTemplate(w, "help", nil)
	} else if r.Method == "POST" {
		redirectHome(w, r)
	} else {
		redirectHome(w, r)
	}
}

/* * */

var sPath string
var templates *template.Template

func renderTemplate(w http.ResponseWriter, template string, res *Result) {
	err := templates.ExecuteTemplate(w, template + ".html", res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

/* * */

func start() {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}
	sPath = path.Dir(filename)

	templates = template.Must(template.ParseFiles(sPath+ "/view/home.html", sPath+ "/view/search.html", sPath+ "/view/help.html"))
}

// RunServer runs the server.
func RunServer() {
	start()

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/search", searchHandler)
	http.HandleFunc("/help", helpHandler)

	server.ListenAndServe()
}