package output

import (
	"ICG_SSAU/generator"
	"ICG_SSAU/tests"
	"fmt"
)

// PrintResults Вывод результатов
func PrintResults(params generator.Params, results []tests.TestResult) {
	fmt.Println("Генерация случайных чисел")
	fmt.Println("Метод: Обратный конгруэнтный генератор")
	fmt.Println("Параметры:")
	fmt.Printf("- a(множитель) = %d\n", params.A)
	fmt.Printf("- c(приращение) = %d\n", params.C)
	fmt.Printf("- n(Модуль) = %d\n", params.N)
	fmt.Printf("- seed = %d\n", params.Seed)
	fmt.Printf("- размер последовательности: %d 000 000\n\n", params.Length/1_000_000)

	fmt.Println("Детальные результаты тестов:")
	allPassed := true
	for _, res := range results {
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
