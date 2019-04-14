// Listing 5.17  Canary test of MyWriter

func TestWriter(t *testing.T) {
	var _ io.Writer = &MyWriter()
}
