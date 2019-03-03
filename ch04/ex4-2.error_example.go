// Listing 4.2  Handling an error

func main() {

	args := os.Args[1:]

	if result, err := Concat(args...); err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Printf("Concatenated string: '%s'\n", result)
	}

}
