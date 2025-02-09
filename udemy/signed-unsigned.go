package main

import (
	"fmt"

	"github.com/kotahemakumargutti/GoLang/TestingModule"
)

func main() {
	var a int8 = 127
	var b uint8 = 255
	fmt.Println(a, b)
	var x, y int64 = 4, 5
	fmt.Println(TestingModule.Addition(x, y))
}
