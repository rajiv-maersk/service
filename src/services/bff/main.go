package main

// add zap logger

import (
	"fmt"

	"github.com/rajiv-maersk/service/bff/helpers"
	"go.uber.org/zap"
)

func main() {
	a, b := 5, 3
	sum := helpers.Add(a, b)
	difference := helpers.Sub(a, b)

	logger, _ := zap.NewProduction()
	defer logger.Sync()
	fmt.Printf("Sum: %d\n", sum)
	fmt.Printf("Difference: %d\n", difference)
}
