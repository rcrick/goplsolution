package main

// go run main.go a s
import (
	"fmt"
	"os"
)

func main() {
	var s, seq string
	for _, arg := range os.Args[1:] {
		s += seq + arg
		seq = " "
	}
	fmt.Println(s)
}
