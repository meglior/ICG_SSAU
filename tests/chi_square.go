package tests

import (
	"fmt"
	"gonum.org/v1/gonum/stat/distuv"
	"math"
)

// ChiSquareTest Тест хи-квадрат
func ChiSquareTest(sequence []int64, m int64, length int64) TestResult {
	// Вычисление количества интервалов(bin) по правилу Стерджесса
	bins := int(math.Ceil(1 + 3.322*math.Log10(float64(length))))

	// Инициализация счетчика для подсчета попаданий в интервал
	counts := make([]int, bins)

	// Распределение значений по интервалам
	for _, num := range sequence {
		bin := num * int64(bins) / m
		if bin >= int64(bins) {
			bin = int64(bins) - 1
		}
		counts[bin]++
	}

	// Вычисление ожидаемого количества значений в каждом интервале
	expected := float64(len(sequence)) / float64(bins)

	// Вычисление статистики хи-квадрат
	var chiSq float64
	for _, count := range counts {
		chiSq += math.Pow(float64(count)-expected, 2) / expected
	}

	// Определение степеней свободы
	df := bins - 1

	// Получение критического значения
	criticalValue := getDynamicCriticalValue(float64(df), 0.05) // Вычисляем динамическое критическое значение для уровня значимости 0.05

	// Проверка условия прохождения теста
	passed := chiSq < criticalValue

	// Возвращение результата теста
	return TestResult{
		Name:        "Хи-квадрат",
		Description: "проверка равномерности распределения",
		Stat:        chiSq,
		Critical:    fmt.Sprintf("%.2f (степени свободы: %d)", criticalValue, df),
		Passed:      passed,
		Metrics: map[string]interface{}{
			"counts": counts,
			"bins":   bins,
		},
	}
}

// getDynamicCriticalValue Динамическое вычисление критического значения для хи-квадрат распределения
func getDynamicCriticalValue(df float64, alpha float64) float64 {
	// Используем формулу для вычисления критического значения
	criticalValue := df*math.Pow(1-2/(9*df)-math.Sqrt(2/(9*df)), 3) + chiSquaredQuantile(alpha, df)
	return criticalValue
}

// chiSquaredQuantile вероятность хи-квадрат распределения
func chiSquaredQuantile(p float64, df float64) float64 {
	// Используем gonum для вычисления вероятностей хи-квадрат распределения
	dist := distuv.ChiSquared{K: df}
	quantile := dist.Quantile(p)
	return quantile
}
