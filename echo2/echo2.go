// Prints its command line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] { // range returns index and element
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
