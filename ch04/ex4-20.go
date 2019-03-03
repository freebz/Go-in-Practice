// Listing 4.20  A send HTTP serve

package main

import (
	"errors"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
