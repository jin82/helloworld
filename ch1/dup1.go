package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var count = make(map[string]int)
	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		if in.Text() == "-c" {
			break
		}
		count[in.Text()]++
	}
	for text, num := range count {
		if num > 1 {
			fmt.Printf("%d\t%s\n", num, text)
		}
	}
}
