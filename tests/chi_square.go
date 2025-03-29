package tests

import (
	"fmt"
	"math"
)

// ChiSquareTest Тест хи-квадрат
func ChiSquareTest(sequence []int64, m int64, bins int) TestResult {
	counts := make([]int, bins)
	for _, num := range sequence {
		bin := num * int64(bins) / m
		if bin >= int64(bins) {
			bin = int64(bins) - 1
		}
		counts[bin]++
	}

	expected := float64(len(sequence)) / float64(bins)
	var chiSq float64
	for _, count := range counts {
		chiSq += math.Pow(float64(count)-expected, 2) / expected
	}

	df := bins - 1
	criticalValue := getChiSquareCriticalValue(df)
	passed := chiSq < criticalValue

	return TestResult{
		Name:        "Хи-квадрат",
		Description: "проверка равномерности распределения",
		Stat:        chiSq,
		Critical:    fmt.Sprintf("%.2f (степени свободы: %d)", criticalValue, df),
		Passed:      passed,
	}
}

// Таблица критических значений хи-квадрат для α=0.05
func getChiSquareCriticalValue(df int) float64 {
	criticalValues := map[int]float64{
		9:   16.919,
		10:  18.307,
		20:  31.410,
		30:  43.773,
		40:  55.758,
		50:  67.505,
		60:  79.082,
		70:  90.531,
		80:  101.879,
		90:  113.145,
		100: 124.342,
	}
	if val, ok := criticalValues[df]; ok {
		return val
	}
	return 0
}
