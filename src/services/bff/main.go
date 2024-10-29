package main

import (
	"fmt"

	"github.com/rajiv-maersk/service/bff/helpers"
	"go.uber.org/zap"
)

// Sub returns the difference of two integers.
func Sub(a, b int) int {
	return a - b
}

// CalculateAndLog performs the addition and subtraction, logging the results.
func CalculateAndLog(a, b int) (int, int, error) {
	sum := helpers.Add(a, b)
	difference := Sub(a, b)

	logger, err := zap.NewProduction()
	if err != nil {
		return 0, 0, err
	}
	defer logger.Sync()

	logger.Info("Calculating values",
		zap.Int("Sum", sum),
		zap.Int("Difference", difference),
	)

	return sum, difference, nil
}

func main() {
	a, b := 5, 3
	sum, difference, err := CalculateAndLog(a, b)
	if err != nil {
		fmt.Printf("Error initializing logger: %v\n", err)
		return
	}
	fmt.Printf("Sum: %d\n", sum)
	fmt.Printf("Difference: %d\n", difference)
}
