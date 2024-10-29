package main

import (
	"fmt"

	"github.com/rajiv-maersk/service/hello-wasm/add"
	"github.com/rajiv-maersk/service/hello-wasm/sub"
	"go.uber.org/zap"
)

func main() {
	a, b := 5, 3
	sum := add.Add(a, b)
	difference := sub.Sub(a, b)

	logger := zap.NewExample()
	logger.Info("Sum", zap.Int("sum", sum))
	fmt.Printf("Sum: %d\n", sum)
	fmt.Printf("Difference: %d\n", difference)
}
