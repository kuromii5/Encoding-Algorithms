package utils

import (
	"fmt"
	"math"
)

func CheckProbs(probabilites []float64) bool {
	sum := 0.0
	for _, num := range probabilites {
		sum += num
	}

	// Define a small value to handle floating-point precision issues
	epsilon := 0.000000001
	if sum >= 1.0-epsilon && sum <= 1.0+epsilon {
		fmt.Println("The sum of floats equals 1")
	} else {
		fmt.Println("The sum of floats does not equal 1")
		return false
	}

	return true
}

func CalculateEntropy(probabilites []float64) float64 {
	entropy := 0.0

	if CheckProbs(probabilites) {
		for _, prob := range probabilites {
			entropy -= prob * math.Log2(prob)
		}
	}

	return entropy
}
