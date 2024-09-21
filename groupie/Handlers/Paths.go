package Handler

import (
	"net/http"
	"text/template"
)

// error messages
const (
	method_err   = "error 405 methode not allowed"
	internal_err = "error 500 Internal Server Error"
	bad_err      = "error 400 BAD REQUEST"
	page_err     = "error 404 page  not found"
)

var (
	artists              []Data
	templates, errT      = template.ParseFiles("Templates/artists.html")
	error_template, errE = template.ParseFiles("Templates/Error.html")
	url                  = "https://groupietrackers.herokuapp.com/api/"
)

func HAndleHOme(w http.ResponseWriter, r *http.Request) {
	CheckparseFIle(w, r, errT, errE)
	if r.Method != "GET" {
		Error_handle(w, 405, method_err)
		return
	}
	if r.URL.Path != "/" {
		Error_handle(w, 400, bad_err)
		return
	}

	// Replace with your actual URL
	if Get_JSONData(url) == nil {
		Error_handle(w, 500, internal_err)
		return
	}

	exec := templates.ExecuteTemplate(w, "artists.html", artists)
	if exec != nil {
	}
}
