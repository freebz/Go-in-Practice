// Listing 7.3  Serving a subdirectory

func main() {
	dir := http.Dir("./files/")
	handler := http.StripPrefix("/static/", http.FileServer(dir))
	http.Handle("/static/", handler)

	http.HandleFunc("/", homePage)
	http.ListenAndServe(":8080", nil)
}
