// Listing 7.7  Embedding files in binaries with go.rice: embedded_files.go

package main

import (
	"github.com/GeertJohan/go.rice"
	"net/http"
)

func main() {
	box := rice.MustFindBox("../files/")
	httpbox := box.HTTPBox()
	http.ListenAndServe(":8080", http.FileServer(httpbox))
}
