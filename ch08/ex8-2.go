// Listing 8.2  DELETE request with default HTTP client

package main

import (
	"fmt"
	"net/http"
)

func main() {
	req, _ := http.NewRequest("DELETE",
		"http://example.com/foo/bar", nil)
	res, _ := http.DefaultClient.Do(req)
	fmt.Printf("%s", res.Status)
}
