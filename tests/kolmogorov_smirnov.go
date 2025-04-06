package tests

import (
	"fmt"  // Для форматированного вывода текста
	"math" // Для математических функций, таких как sqrt и Abs
	"sort" // Для сортировки массива с плавающей точкой
)

// KolmogorovSmirnovTest Функция для проведения теста Колмогорова-Смирнова
func KolmogorovSmirnovTest(sequence []int64, n int64) TestResult {
	l := len(sequence) // Длина входной последовательности
	// создаем копию последовательности, нормированную к интервалу [0, 1]
	sorted := make([]float64, l)
	for i, num := range sequence {
		sorted[i] = float64(num) / float64(n) // Нормализация значений
	}
	// Сортируем нормализованную последовательность
	sort.Float64s(sorted)

	d := 0.0 // Инициализация переменных для расчета статистики D
	for i := 0; i < l; i++ {
		observed := sorted[i]                     // Наблюдаемое значение
		expected := float64(i+1) / float64(l)     // Ожидаемое значение (для равномерного распределения)
		currentD := math.Abs(expected - observed) // Рассчитываем разницу между наблюдаемым и ожидаемым значением
		if currentD > d {
			d = currentD // Сохраняем максимальную разницу среди всех точек
		}
	}

	criticalValue := 1.36 / math.Sqrt(float64(l)) // Расчет критического значения для уровня значимости 0.05
	passed := d < criticalValue                   // Проверка, прошла ли последовательность тест

	return TestResult{
		Name:        "Колмогорова-Смирнова",                 // Название теста
		Description: "проверка равномерности распределения", // Описание теста
		Stat:        d,                                      // Значение статистики D
		Critical:    fmt.Sprintf("%.4f", criticalValue),     // Критическое значение
		Passed:      passed,                                 // Флаг прохождения теста
	}
}
