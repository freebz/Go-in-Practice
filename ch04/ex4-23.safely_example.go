// Listing 4.23  Using safely.Go to trap panics

package main

import (
	"github.com/Masterminds/cookoo/safely"
	"errors"
	"time"
)

func message() {
	println("Inside goroutine")
	panic(errors.New("Oops!"))
}

func main() {
	safely.Go(message)
	println("Outside goroutine")
	time.Sleep(1000)
}
