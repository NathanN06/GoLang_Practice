package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {

	//Define text
	text := r.URL.Query().Get("text")

	//Create if statement to handle nothing sent
	if text == "" {
		text = "stranger"
	}

	//Print line
	fmt.Fprintf(w, "Hello, %s", text)
}

func main() {

	//Define path
	http.HandleFunc("/hello", helloHandler)

	//Print helper
	fmt.Println("Server running at http://localhost:8080/")

	//Activate server
	http.ListenAndServe(":8080", nil)

}
