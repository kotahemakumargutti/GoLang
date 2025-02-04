package main

import "fmt"

type KiloByte int

const (
	_ KiloByte = 1 << (10 * iota)
	KB
	MB
	GB
)

func main() {
	fmt.Printf("%b\t%d\n", KB, KB)
	fmt.Printf("%b\t%d\n", MB, MB)
	fmt.Printf("%b\t%d\n", GB, GB)
}
