// Listing 4.4  Relying on good error handling

func main() {
	args := os.Args[1:]
	result, _ := Concat(args...)
	fmt.Printf("Concatenated string: '%s'\n", result)
}
