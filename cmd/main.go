package main

import (
	"ICG_SSAU/generator" // Импортируем пакет генератора псевдослучайных чисел
	"ICG_SSAU/output"    // Импортируем пакет вывода результатов
	"ICG_SSAU/tests"     // Импортируем пакет тестирования последовательности
	"time"               // Пакет для работы с временем, необходим для генерации начального значения
)

func main() {
	// Параметры генератора (подробное описание в generator/generator.go)
	seed := time.Now().UnixNano() % 2147483647 // Генерируем начальное значение на основе текущего времени
	a := int64(1103515245)                     // Множитель a (параметр линейного конгруэнтного генератора) (0 ≤ a < n)
	c := int64(12345)                          // Приращение c (параметр линейного конгруэнтного генератора) (0 ≤ c < n)
	n := int64(2147483647)                     // Модуль n (определяющий пространство значений) (2^64 - 1, простое число)
	length := int64(1000)                      // Длина генерируемой последовательности

	// Генерация последовательности
	sequence, params := generator.GenerateInverseLCG(seed, a, c, n, length)

	// Параллельный запуск тестов (подробности в tests/*.go)
	results := make(chan tests.TestResult, 4)

	go func() { results <- tests.ChiSquareTest(sequence, n, length) }() // Тест хи-квадрат
	go func() { results <- tests.KolmogorovSmirnovTest(sequence, n) }() // Критерий Колмогорова-Смирнова
	go func() { results <- tests.CramerVonMisesTest(sequence, n) }()    // Критерий Крамер-фон Мизеса
	go func() { results <- tests.SpectralTest(sequence) }()             // Спектральный тест
	go func() { results <- tests.DurbinWatsonTest(sequence) }()         // Тест Дарбина-Уотсона
	// Сбор результатов и вывод
	var testResults []tests.TestResult
	for i := 0; i < 5; i++ {
		testResults = append(testResults, <-results) // Добавляем результаты тестов в список
	}
	// Печать результатов через модуль output
	output.PrintResults(params, testResults)
}
