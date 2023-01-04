package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "GET" {

			host := r.Host

			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "Hello, %s", host)

		}
	})

	fmt.Println("Server is running on port 4000")
	http.ListenAndServe(":4000", nil)

}
