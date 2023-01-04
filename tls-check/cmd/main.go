package main

import (
	"fmt"
	"net/http"
)

func main() {

	allowedDomains := []string{"shrti.xyz", "shrti.io"}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "GET" {

			query := r.URL.Query()

			requestDomain := query.Get("domain")

			if requestDomain == "" {
				w.WriteHeader(http.StatusNotFound)
				return
			}

			for _, domain := range allowedDomains {
				if requestDomain == domain {
					w.WriteHeader(http.StatusOK)
					fmt.Fprintf(w, "OK")

					fmt.Println("Domain is allowed: ", requestDomain)

					return
				}
			}

			fmt.Println("Domain is not allowed: ", requestDomain)

			w.WriteHeader(http.StatusNotFound)
			return

		}
	})

	fmt.Println("TLS check server is running on port 5555")
	http.ListenAndServe(":5555", nil)

}
