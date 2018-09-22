// Prints its command line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""                  // Short variable declaration can only be used inside a function.
	for _, arg := range os.Args[1:] { // range returns index and element, _ is used to discard unused values.
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
