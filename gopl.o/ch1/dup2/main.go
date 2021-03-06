package main

import (
	"bufio"
	"fmt"
	"os"
)

// go run main.go 1.txt ss
func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			}
			defer f.Close()
			countLines(f, counts)
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Println(line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
