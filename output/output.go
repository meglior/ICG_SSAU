package output

import (
	"ICG_SSAU/generator"
	"ICG_SSAU/tests"
	"fmt"
	"github.com/guptarohit/asciigraph"
)

// PrintResults Вывод результатов
func PrintResults(params generator.Params, results []tests.TestResult) {
	fmt.Println("Генерация случайных чисел")
	fmt.Println("Метод: Обратный конгруэнтный генератор")
	fmt.Println("Параметры:")
	fmt.Printf("- a(множитель) = %d\n", params.A)
	fmt.Printf("- c(приращение) = %d\n", params.C)
	fmt.Printf("- n(Модуль) = %d\n", params.N)
	fmt.Printf("- seed(начальное значение) = %d\n", params.Seed)
	fmt.Printf("- размер последовательности: %d 000 000\n\n", params.Length/1_000_000)

	fmt.Println("Детальные результаты тестов:")
	allPassed := true
	for _, res := range results {
		if res.Name == "Хи-квадрат" {
			if counts, ok := res.Metrics["counts"].([]int); ok {
				if bins, ok2 := res.Metrics["bins"].(int); ok2 {
					plotHistogram(counts, bins)
				}
			}
		}
		fmt.Printf("\nТест %s: %s\n", res.Name, res.Description)
		if res.Passed {
			fmt.Println("✅ ПРОЙДЕН")
		} else {
			fmt.Println("❌ НЕ ПРОЙДЕН")
			allPassed = false
		}
		fmt.Printf("Статистика: %.4f\n", res.Stat)
		fmt.Printf("Критическое значение: %s\n", res.Critical)
	}

	fmt.Println("\nИтог:")
	if allPassed {
		fmt.Println("Все тесты пройдены успешно. Последовательность соответствует требованиям случайности.")
	} else {
		fmt.Println("НЕ ВСЕ ТЕСТЫ ПРОЙДЕНЫ! Последовательность не удовлетворяет критериям случайности.")
	}
}

func plotHistogram(counts []int, bins int) {
	data := make([]float64, len(counts))
	for i, c := range counts {
		data[i] = float64(c)
	}

	graph := asciigraph.Plot(
		data,
		asciigraph.Height(20),
		asciigraph.Width(70),
		//asciigraph.Caption("Распределение значений по интервалам"),
	)

	fmt.Println("\nГистограмма распределения:")
	fmt.Println(graph)
}
