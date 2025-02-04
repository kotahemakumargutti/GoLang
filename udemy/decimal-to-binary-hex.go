package main

import "fmt"

func main() {
	a, b, c := 747, 911, 90210
	fmt.Printf("%d in hexa %X in binary %b\n", a, a, a)
	fmt.Printf("%d in hexa %X in binary %b\n", b, b, b)
	fmt.Printf("%d in hexa %#X in binary %b\n", c, c, c)

	x, _ := fmt.Printf("%X", a)
	fmt.Printf("%T\n", x)
}
