// Listing 5.11  The stringer interface

type Writer interface {
	Write(p []byte) (n int, err error)
}
