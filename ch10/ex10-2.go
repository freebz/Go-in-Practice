// Listing 10.2  Falling to close an HTTP response body

func main() {
	r, err := http.Get("http://example.com")
	if err != nil {
		...
	}
	defer r.Body.Close()
	o, err := ioutil.ReadAll(r.Body)
	if err != nil {
		...
	}
	// Use the body content

	r2, err := http.Get("http://example.com/foo")
	if err != nil {
		...
	}
	defer r2.Body.Close()
	o, err = ioutil.ReadAll(r2.Body)
	if err != nil {
		...
	}
	...
}
