// Listing 7.11  Parsing a form multiple values for a field

func exampleHandler(w http.ResponseWriter, r *http.Request) {
	maxMemory := 16 << 20
	err := r.ParseMultipartForm(maxMemory)
	if err != nil {
		fmt.Println(err)
	}
	for k, v := range r.PostForm["names"] {
		fmt.Println(v)
	}
}
