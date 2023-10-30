package main

import "fmt"

const x int64 = 10

const (
	idkey   = "id"
	nameKey = "name"
)

const z = 20 * 10

func main() {
	const y = "hello"

	fmt.Print(x)
	fmt.Print(y)

	x = x + 1
	y = "bye"

	fmt.Print(x)
	fmt.Print(y)
}
