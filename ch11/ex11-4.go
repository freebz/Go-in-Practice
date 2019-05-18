// Listing 11.4  Checking and converting a type

package main

import (
	"bytes"
	"fmt"
)

func main() {
	b := bytes.NewBuffer([]byte("Hello"))
	if isStringer(b) {
		fmt.Printf("%T is a stringer\n", b)
	}
	i := 123
	if isStringer(i) {
		fmt.Printf("%T is a stringer\n", i)
	}
}

func isStringer(v interface{}) bool {
	_, ok := v.(fmt.Stringer)
	return ok
}
