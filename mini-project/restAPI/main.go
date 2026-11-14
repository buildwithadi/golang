package main

import "net/http"

func homeHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	http.HandleFunc("/", homeHandler)

	http.ListenAndServe(":8000", nil)
}
