// Listing 7.2  Serve file with custom handler: servefile.go

package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", readme)
	http.ListenAndServe(":8080", nil)
}

func readme(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "./files/readme.txt")
}
