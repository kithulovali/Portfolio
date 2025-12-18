package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "the server is running okay")
}
func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":8081", nil)
}
