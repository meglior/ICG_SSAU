package tests

import (
	"math"
	"sort"
)

// CramerVonMisesTest Тест Крамера-фон Мизеса
func CramerVonMisesTest(sequence []int64, m int64) TestResult {
	n := len(sequence)
	sorted := make([]float64, n)
	for i, num := range sequence {
		sorted[i] = float64(num) / float64(m)
	}
	sort.Float64s(sorted)

	var w2 float64
	for i := 0; i < n; i++ {
		u := sorted[i]
		w2 += math.Pow(u-(float64(2*i+1)/(2*float64(n))), 2)
	}
	w2 = (1.0 / (12.0 * float64(n))) + w2

	passed := w2 < 0.461 // Критическое значение для α=0.05
	return TestResult{
		Name:        "Крамера-фон Мизеса",
		Description: "проверка равномерности распределения",
		Stat:        w2,
		Critical:    "0.461",
		Passed:      passed,
	}
}
