package main

import (
	"fmt"

	"github.com/rajiv-maersk/service/hello-wasm/helpers"
)

func main() {
	a, b := 5, 3
	sum := helpers.Add(a, b)
	difference := helpers.Sub(a, b)

	fmt.Printf("Sum: %d\n", sum)
	fmt.Printf("Difference: %d\n", difference)
}
