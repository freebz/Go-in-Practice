// Listing 7.10  Parsing a simple form response

func exampleHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
	}
	name := r.FormValue("name")
}
