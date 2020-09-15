package main

import "fmt"

func zeroval(i int) {
	i = 0
}

func zeroptr(ptr *int) {
	*ptr = 0
}

func main() {
	i := 1
	fmt.Println("initial: ", i)

	zeroval(i)
	fmt.Println("no pointer: ", i)

	zeroptr(&i)
	fmt.Println("with pointer: ", i)
}
