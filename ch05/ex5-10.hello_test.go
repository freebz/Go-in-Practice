// Listing 5.10  A hello test

package hello

import "testing"

func TestHello(t *testing.T) {
	if v := Hello(); v != "hello" {
		t.Errorf("Expected 'hello', but got '%s'", v)
	}
}
