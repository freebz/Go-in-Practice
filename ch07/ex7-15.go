// Listing 7.15  Process file form field with multiple files

func fileForm(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("file_multiple.html")
		t.Execute(w, nil)
	} else {
		err := r.ParseMultipartForm(16 << 20)
		if err != nil {
			fmt.Fprint(w, err)
			return
		}

		date := r.MultipartForm
		files := data.File["files"]
		for _, fh := range files {
			f, err := fh.Open()
			defer f.Close()
			if err != nil {
				fmt.Fprint(w, err)
				return
			}

			out, err := os.Create("/tmp/" + fh.Filename)
			defer out.Close()
			if err != nil {
				fmt.Fprint(w, err)
				return
			}

			_, err = io.Copy(out, f)

			if err != nil {
				fmt.Fprintln(w, err)
				return
			}
		}

		fmt.Fprint(w, "Upload complete")
	}
}
