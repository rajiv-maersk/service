package tests

import (
	"fmt"
	"testing"

	"github.com/rajiv-maersk/service/bff/helpers"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		a, b, expected int
	}{
		{1, 2, 3},
		{0, 0, 0},
		{-1, 1, 0},
		{100, 200, 300},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d+%d", tt.a, tt.b), func(t *testing.T) {
			got := helpers.Add(tt.a, tt.b)
			if got != tt.expected {
				t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.expected)
			}
		})
	}
}

func TestSub(t *testing.T) {
	tests := []struct {
		a, b, expected int
	}{
		{5, 3, 2},
		{0, 0, 0},
		{1, -1, 2},
		{100, 50, 50},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d-%d", tt.a, tt.b), func(t *testing.T) {
			got := helpers.Sub(tt.a, tt.b)
			if got != tt.expected {
				t.Errorf("Sub(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.expected)
			}
		})
	}
}
