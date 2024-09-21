package Handler

import (
	"net/http"
)

// function to execute the errors
func Error_handle(w http.ResponseWriter, num int, message string) {
	w.WriteHeader(num)
	error_template.Execute(w , message )
	
}

// function to check if there's a problem in parsing error and index html files
func CheckparseFIle(w http.ResponseWriter, r *http.Request, err1 error, err2 error)  {
	if err2 != nil {
		http.Error(w, "Internal Server Error () ", 500) 
		return
	} else if err1 != nil || err2 != nil {
		Error_handle(w, 500, "error 500 Internal Server Error")
	}
}
