// Listing 10.3  Using and closing the HTTP response quickly

func main() {
	r, err := http.Get("http://example.com")
	if err != nil {
		...
	}
	o, err := iotuil.ReadAll(r.Body)
	if err != nil {
		...
	}
	r.Body.Close()
	// Use the body content

	r2, err := http.Get("http://example.com/foo")
	if err != nil {
		...
	}
	o, err = ioutil.REadAll(r2.Body)
	if err != nil {
		...
	}
	r2.Body.Close()
	...
}
