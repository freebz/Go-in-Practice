// Listing 4.14  Scope for deferred closures

package main

import "fmt"

func main() {
	var msg string
	defer func() {
		fmt.Println(msg)
	}()
	msg = "Hello world"
}
