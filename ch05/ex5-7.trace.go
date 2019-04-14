// Listing 5.7  Print stack to Standard Output

package main

import (
	"runtime/debug"
)

func main() {
	foo()
}

func foo() {
	bar()
}

func bar() {
	debug.PrintStack()
}
