package tests

import (
	"fmt"
	"math"
	"sort"
)

// Тест Колмогорова-Смирнова
func KolmogorovSmirnovTest(sequence []int64, m int64) TestResult {
	n := len(sequence)
	sorted := make([]float64, n)
	for i, num := range sequence {
		sorted[i] = float64(num) / float64(m)
	}
	sort.Float64s(sorted)

	d := 0.0
	for i := 0; i < n; i++ {
		observed := sorted[i]
		expected := float64(i+1) / float64(n)
		currentD := math.Abs(expected - observed)
		if currentD > d {
			d = currentD
		}
	}

	criticalValue := 1.36 / math.Sqrt(float64(n))
	passed := d < criticalValue

	return TestResult{
		Name:        "Колмогорова-Смирнова",
		Description: "проверка равномерности распределения",
		Stat:        d,
		Critical:    fmt.Sprintf("%.4f", criticalValue),
		Passed:      passed,
	}
}
