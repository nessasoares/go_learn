package main

import "fmt"

func main() {
	i := 0
HERE:
	fmt.Printf("%d", i)
	i++
	if i == 5 {
		return
	}

	goto HERE
}
