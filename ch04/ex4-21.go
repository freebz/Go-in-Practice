// Listing 4.21  A panicky handler

function handler(res http.ResponseWriter, req *http.Request) {
	panic(errors.New("Fake panic!"))
}
