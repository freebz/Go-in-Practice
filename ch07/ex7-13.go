// Listing 7.13  Handle a single file upload

func fileForm(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("file.html")
		t.Execute(w, nil)
	} else {
		f, h, err := r.FromFile("file")
		if err != nil {
			panic(err)
		}
		defer f.Close()
		filename := "/tmp/" + h.Filename
		out, err := os.Create(filename)
		if err != nil {
			panic(err)
		}
		defer out.Close()
		io.Copy(out, f)
		fmt.Fprint(w, "Upload complete")
	}
}
