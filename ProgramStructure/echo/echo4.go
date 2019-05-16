// Echo4 prints its command-line arguments
package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline") // n and sep are pointers to the flag variables
var sep = flag.String("s", " ", "separator")           // and must be accessed indirectly as *n and *sep

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
