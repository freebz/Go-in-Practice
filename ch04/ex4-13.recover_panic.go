// Listing 4.13  Recovering from a panic

package main

import (
	"errors"
	"fmt"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Trapped panic: % (%T)\n", err, err)
		}
	}()

	yikes()
}

func yikes() {
	panic(errors.New("Something bad happened."))
}
