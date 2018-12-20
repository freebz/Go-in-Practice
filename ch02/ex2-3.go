// Listing 2.3  Custom flag help text

flag.VisitAll(func(flag * flag.Flag) {
	format := "\t-%s: %s (Default: '%s')\n"
	fmt.Printf(format, flag.Name, flag.Usage, flag.DefValue)
})
