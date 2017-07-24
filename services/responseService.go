package services

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ErrorWithJSON(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	fmt.Println("{message: %q}", message)
	/*_, err := fmt.Fprintf(w, "{message: %q}", message)
	if err != nil {
		fmt.Println(err)
	}*/
	var err struct{ Errors string }
	err.Errors = message

	resp, error := json.Marshal(err)
	if error != nil {
		fmt.Println("Error Response errored out")
	}

	w.Write(resp)
}

func ResponseWithJSON(w http.ResponseWriter, json []byte, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	w.Write(json)
	return
}
