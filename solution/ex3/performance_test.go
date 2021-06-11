// go test -bench=. 
package test

import (
	"strings"
	"testing"
)

var args = []string{"hi", "there", "buddy", "boy", "5", "6", "7", "8", "9"}

func BenchmarkContact(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var s, seq string
		for _, arg := range args {
			s += seq + arg
			seq = " "
		}
	}
}

func BenchmarkJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.Join(args, " ")
	}
}
