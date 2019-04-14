// Listing 5.14  Testing with a mock

package msg

import (
	"testing"
)

type MockMessage struct {
	email, subject string
	body           []byte
}

func (m *MockMessage) Send(email, subject string, body []byte) error {
	m.email = email
	m.subject = subject
	m.body = body
	return nil
}

func TestAlert(t *testing.T) {
	msgr := new(MockMessage)
	body := []byte("Critical Error")

	Alert(msgr, body)

	if msgr.subject != "Critical Error" {
		t.Errorf("Expected 'Critial Error', Got '%s'", msgr.subject)
	}
	// ...
}
