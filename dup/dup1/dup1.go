// Prints  the text of each line that appears more than
// once in the standard input, preceded by its count.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)      // Map holds key/value pairs with O(1) time for storing, retrieving and testing for an item in the set.
	input := bufio.NewScanner(os.Stdin) // Scanner reads input and breaks it into lines or words.
	for input.Scan() {                  // Reads the next line and removes the newline character from the end.
		// Returns true if there is a line and false otherwise.
		counts[input.Text()]++ // The first time a new line is seen, counts[input.Text()] evaluates to the zero value for its type (0 for int).
	}
	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
