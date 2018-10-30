// Listing 1.4  HTTP GET: http_get.go

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, _ :=
		http.Get("http://example.com")
	body, _ :=
		ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	resp.Body.Close()
}
