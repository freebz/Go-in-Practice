// Listing 7.4  Using path package path resolution

func main() {
	pr := newPathResolver()
	pr.Add("GET /hello", hello)

	dir := http.Dir("./files")
	handler := http.StripPrefix("/static/", http.FileServer(dir))
	pr.Add("GET /static/*", handler.ServeHTTP)

	http.ListenAndServe(":8080", pr)
}
