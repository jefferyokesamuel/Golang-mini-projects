package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	counter := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := os.ReadFile(filename)
		if err != nil { 
			fmt.Fprintf(os.Stderr, "dup3S: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counter[line]++
		}
	}

	for line, n := range counter { 
		if n > 1 {
			fmt.Printf("%d/t%s\n", n, line)
		}
	}
}