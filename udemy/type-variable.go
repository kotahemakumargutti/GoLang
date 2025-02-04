package main

import "fmt"

func main() {
	var a string = "Hello World"
	var b = 10
	c := 20.0
	fmt.Printf("%v of type %T\n", a, a)
	fmt.Printf("%v of type %T\n", b, b)
	fmt.Printf("%v of type %T\n", c, c)
}
