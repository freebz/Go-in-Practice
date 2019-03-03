// Listing 4.3  Returning useful data with an error

func Concat(parts ...string) (string, error) {
	if len(parts) == 0 {
		return "", errors.New("No strings supplied")
	}

	return strings.Join(parts, " "), nil
}
