// Prints the count and text of lines that appear more than once in
// the input. It reads from stdin or from a list of named files.
// For files, prints the names of all the files in shich each duplicated line occurs.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	dupingFiles := make([]bool, len(os.Args))
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for n, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			if countLines(f, counts) {
				dupingFiles[n] = true
			}
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
	for n, arg := range files {
		if dupingFiles[n] {
			fmt.Println(arg)
		}
	}
}

func countLines(f *os.File, counts map[string]int) bool { // Functions and other package-level entities may be declared in any order.
	input := bufio.NewScanner(f)
	dupCheck := false
	for input.Scan() {
		counts[input.Text()]++
		if counts[input.Text()] > 1 {
			dupCheck = true
		}
	}
	// NOTE: ignoring potential errors from input.Err()
	return dupCheck
}
