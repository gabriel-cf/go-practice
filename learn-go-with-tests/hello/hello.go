package main

import(
	"fmt"
)

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum - 1
	y = sum - 2
	return
}

const hello = "Hello "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return hello + name
}

func main() {
	fmt.Printf(Hello("World"))
}