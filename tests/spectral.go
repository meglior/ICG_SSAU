package tests

import (
	"fmt"
	"math"
	"math/cmplx"
)

// SpectralTest реализует спектральный тест (NIST STS)
func SpectralTest(sequence []int64) TestResult {
	// 1. Преобразуем числа в битовую последовательность
	bits := toBitSequence(sequence)
	n := len(bits)

	// 2. Проверяем минимальную длину последовательности
	if n < 1000 {
		return TestResult{
			Name:        "Спектральный тест",
			Description: "проверка многомерной равномерности",
			Passed:      false,
			Critical:    "Слишком короткая последовательность (минимум 1000 бит)",
		}
	}

	// 3. Преобразование в комплексный спектр
	spectrum := make([]complex128, n)
	for i, bit := range bits {
		spectrum[i] = complex(float64(bit), 0)
	}

	// 4. Выполняем быстрое преобразование Фурье
	fft(spectrum)

	// 5. Вычисляем модули спектральных компонент
	magnitudes := make([]float64, n/2)
	for i := 0; i < n/2; i++ {
		magnitudes[i] = cmplx.Abs(spectrum[i])
	}

	// 6. Находим пики в спектре
	peakCount := countSignificantPeaks(magnitudes, n)

	// 7. Сравниваем с ожидаемым распределением
	expectedPeaks := 0.95 * float64(n) / 2
	zScore := (float64(peakCount) - expectedPeaks) / math.Sqrt(expectedPeaks)
	pValue := math.Erfc(math.Abs(zScore) / math.Sqrt(2))

	// 8. Формируем результат
	passed := pValue >= 0.01
	return TestResult{
		Name:        "Спектральный тест",
		Description: "проверка многомерной равномерности",
		Stat:        zScore,
		Critical:    fmt.Sprintf("p-value ≥ 0.01 (получено: %.4f)", pValue),
		Passed:      passed,
	}
}

// Быстрое преобразование Фурье (рекурсивная реализация)
func fft(x []complex128) {
	n := len(x)
	if n <= 1 {
		return
	}

	// Разделяем на четные и нечетные индексы
	even := make([]complex128, n/2)
	odd := make([]complex128, n/2)
	for i := 0; i < n/2; i++ {
		even[i] = x[2*i]
		odd[i] = x[2*i+1]
	}

	// Рекурсивные вызовы
	fft(even)
	fft(odd)

	// Объединение результатов
	for k := 0; k < n/2; k++ {
		t := cmplx.Exp(complex(0, -2*math.Pi*float64(k)/float64(n))) * odd[k]
		x[k] = even[k] + t
		x[k+n/2] = even[k] - t
	}
}

// Подсчет значимых пиков
func countSignificantPeaks(magnitudes []float64, n int) int {
	threshold := 2.0 / math.Sqrt(float64(n))
	count := 0
	for i := 1; i < len(magnitudes)-1; i++ {
		if magnitudes[i] > magnitudes[i-1] &&
			magnitudes[i] > magnitudes[i+1] &&
			magnitudes[i] > threshold {
			count++
		}
	}
	return count
}

// Вспомогательная функция для преобразования чисел в биты
func toBitSequence(numbers []int64) []int {
	bits := make([]int, 0)
	for _, num := range numbers {
		for i := 0; i < 32; i++ {
			bits = append(bits, int((num>>i)&1))
		}
	}
	return bits
}
