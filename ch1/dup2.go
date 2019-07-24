package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) <= 1 {
		fmt.Println("No Args")
	} else {
		counts := make(map[string]int)
		for _, arg := range os.Args[1:] {
			file, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "open file &s error!\n", arg)
				fmt.Fprintln(os.Stderr, err)
				continue
			}
			countLines(file, counts)
			file.Close()
		}
		for str, num := range counts {
			if num > 1 {
				fmt.Printf("%s\t%d\n", str, num)
			}
		}
	}

}

func countLines(file *os.File, counts map[string]int) {
	input := bufio.NewScanner(file)
	for input.Scan() {
		counts[input.Text()]++
	}
}
