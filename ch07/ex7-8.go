// Listing 7.8  Templates as embedded files

box := rice.MustFindBox("templates")
templateString, err := box.String("example.html")
if err != nil {
	log.Fatal(err)
}
