// Listing 1.11  Hello World test: hello_test.go

package main

import "testing"

func TestName(t *testing.T) {
	name := getName()

	if name != "World!" {
		t.Error("Respone from getName is unexpected value")
	}
}
