// Listing 4.15  msg out of scope

package main

import "fmt"

func main() {
	defer func() {
		fmt.Println(msg)
	}()
	msg := "Hello world"
}
