// Listing 11.15  The Unmarshal function

func Unmarshal(data []byte, v interface{}) error {

	val := reflect.Indirect(reflect.ValueOf(v))
	t := val.Type()

	b := bytes.NewBuffer(data)
	scanner := bufio.NewScanner(b)
	for scanner.Scan() {
		line := scanner.Text()
		pair := strings.SplitN(line, "=", 2)
		if len(pair) < 2 {
			// Skip any malformed lines.
			continue
		}
		setField(pair[0], pair[1], t, val)
	}
	return nil
}
