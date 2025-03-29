package tests

import (
	"math"
)

// DurbinWatsonTest Тест Дарбина-Уотсона
func DurbinWatsonTest(sequence []int64) TestResult {
	sumSquares := 0.0
	sumDiffSquares := 0.0

	for i := 1; i < len(sequence); i++ {
		diff := float64(sequence[i] - sequence[i-1])
		sumDiffSquares += diff * diff
		sumSquares += float64(sequence[i]) * float64(sequence[i])
	}

	dw := sumDiffSquares / sumSquares
	passed := math.Abs(dw-2) < 0.04 // Диапазон [1.96; 2.04]

	return TestResult{
		Name:        "Дарбина-Уотсона",
		Description: "проверка автокорреляции",
		Stat:        dw,
		Critical:    "[1.96 - 2.04]",
		Passed:      passed,
	}
}
