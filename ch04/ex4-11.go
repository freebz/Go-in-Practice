// Listing 4.11  A proper panic

package main

import "errors"

func main() {
	panic(errors.New("Something bad happened."))
}
