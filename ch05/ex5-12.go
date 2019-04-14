// Listing 5.12  The message struct

type Message struct {
	// ...
}

func (m *Message) Send(email, subject string, body []byte) error {
	// ...
	return nil
}
