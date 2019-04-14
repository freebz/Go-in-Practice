// Listing 5.13  use an interface

type Messager interface {
	Send(email, subject string, body []byte) error
}

func Alert(m Messager, problem []byte) error {
	return m.Send("noc@example.com", "Critical Error", problem)
}
