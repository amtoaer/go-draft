package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// string->(filename->count)
	counts := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: %v", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, fileNames := range counts {
		fmt.Printf("[found in %d files] %s\n", len(fileNames), line)
		for fileName, count := range fileNames {
			fmt.Printf("	%d hits in %s\n", count, fileName)
		}
	}
}
func countLines(f *os.File, m map[string]map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		str := input.Text()
		if m[str] == nil {
			m[str] = make(map[string]int)
		}
		m[str][f.Name()]++
	}
}
