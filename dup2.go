// dup2 program from The Go Programming Language
// Exercise 1.4 + do not count empty lines
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	fileCount := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, fileCount)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, fileCount)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
	// print files with double lines and total count
	for file, n := range fileCount {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, file)
		}
	}
}

func countLines(f *os.File, counts map[string]int, fileCount map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		txt := input.Text()
		// do not count empty lines
		if txt != "" {
			counts[txt]++
			// fileCounts remembers the file name, if a line is double one
			// the file that contains the first line is not remembered
			if counts[txt] > 1 {
				fileCount[f.Name()]++
			}
		}
	}
}
