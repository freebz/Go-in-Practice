// Listing 3.2  An anonymous goroutine

package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Outside a goroutine.")
	go func() {
		fmt.Println("Inside a goroutine")
	}()
	fmt.Println("Outside again.")

	runtime.Gosched()
}
