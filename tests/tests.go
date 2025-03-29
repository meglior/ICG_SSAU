package tests

import "fmt"

// TestResult описывает результат отдельного теста
type TestResult struct {
	Name        string
	Description string
	Stat        float64
	Critical    string
	Passed      bool
	Metrics     map[string]interface{}
}

// Форматированный вывод результата теста
func (tr TestResult) String() string {
	status := "✅ ПРОЙДЕН"
	if !tr.Passed {
		status = "❌ НЕ ПРОЙДЕН"
	}
	return fmt.Sprintf("%s: Статистика=%.4f Критическое значение=%s",
		status, tr.Stat, tr.Critical)
}
