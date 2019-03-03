// Listing 4.18  Panic in the response

func response(data []byte, conn net.Conn) {
	panic(errors.New("Failure in response!"))
}
