package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Feels good to B inside this docker container")

}

func main() {
	fmt.Println("I am running on port 8081...")
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8081", nil)
}
