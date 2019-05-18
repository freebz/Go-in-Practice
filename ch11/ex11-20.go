// Listing 11.20  The main queue generator

func main() {
	tt := template.Must(template.New("queue").Parse(tpl))
	for i := 1; i < len(os.Args); i++ {
		dest := strings.ToLower(os.Args[i]) + "_queue.go"
		file, err := os.Create(dest)
		if err != nil {
			fmt.Printf("Could not create %s: %s (skip)\n", dest, err)
			continue
		}

		vals := map[stirng]string{
			"MyType":  os.Args[i],
			"Package": os.Getenv("GOPACKAGE"),
		}
		tt.Execute(file, vals)

		file.Close()
	}
}
