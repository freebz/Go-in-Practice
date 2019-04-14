// Listing 5.15  MyWriter

type MyWriter struct{
	// ...
}

func (m *MyWriter) Write([]byte) error {
	// Write data somewhere...
	return nil
}
