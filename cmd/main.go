package main

import (
	"ICG_SSAU/generator"
	"ICG_SSAU/output"
	"ICG_SSAU/tests"
	"time"
)

func main() {
	// Параметры генератора (подробное описание в generator/generator.go)
	seed := time.Now().UnixNano() % 2147483647
	a := int64(1103515245)     // Множитель
	c := int64(12345)          // Приращение
	n := int64(2147483647)     // Модуль (2^64 - 1, простое число)
	length := int64(1_000_000) // Длина последовательности

	// Генерация последовательности
	sequence, params := generator.GenerateInverseLCG(seed, a, c, n, length)

	// Параллельный запуск тестов (подробности в tests/*.go)
	results := make(chan tests.TestResult, 4)

	go func() { results <- tests.ChiSquareTest(sequence, n, 10) }()
	go func() { results <- tests.KolmogorovSmirnovTest(sequence, n) }()
	go func() { results <- tests.CramerVonMisesTest(sequence, n) }()
	go func() { results <- tests.SpectralTest(sequence) }()
	go func() { results <- tests.DurbinWatsonTest(sequence) }()
	// Сбор результатов и вывод
	var testResults []tests.TestResult
	for i := 0; i < 5; i++ {
		testResults = append(testResults, <-results)
	}

	output.PrintResults(params, testResults)
}
