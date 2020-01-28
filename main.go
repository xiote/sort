package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	names := []string{}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var name string = scanner.Text()

		if len(name) == 0 {
			sorted := Sort(names)
			fmt.Printf("len=%d cap=%d %v\n", len(sorted), cap(sorted), sorted)
			break
		}

		names = append(names, name)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
