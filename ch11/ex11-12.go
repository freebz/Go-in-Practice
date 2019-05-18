// Listing 11.12  The marshal and unmarshal pattern

func Marshal(v interface{}) ([]byte, error) {}
func Unmarshal(data []byte, v interface{}) error {}
