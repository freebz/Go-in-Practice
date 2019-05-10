// Listing 7.1  http package file serving: file_serving.go

package main

import (
	"net/http"
)

func main() {
	dir := http.Dir("./files")
	http.ListenAndServe(":8080", http.FileServer(dir))
}
