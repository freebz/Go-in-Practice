// Listing 3.4  Compressing files in parallel with a wait group

package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var i int = -1
	var file string
	for i, file = range os.Args[1:] {
		wg.Add(1)
		go func(filename string) {
			compress(filename)
			wg.Done()
		}(file)
	}
	wg.Wait()
	
	fmt.Printf("Compressed %d files\n", i+1)
}

func compress(filename string) error {
	// Unchanged from above
}
