package main

import (
	"fmt"
	"net/http"
)

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var text string
	text = "Welcome to my page!"
	w.Write([]byte(text))
}

func main() {
	http.HandleFunc("/", ServeHTTP)
	fmt.Println("Web Server Start!")
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		fmt.Println("Error!")
	}
}
