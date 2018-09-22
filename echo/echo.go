// Print command line arguments.
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args, " ")) // Also print the name of the command that invoked it.
	for i, s := range os.Args[1:] {
		fmt.Println(strconv.Itoa(i) + " " + s) // strconv.Itoa converts integers to string.
	}
}
