package generator

import "math/big"

// Params GeneratorParams Параметры генератора
type Params struct {
	A      int64
	C      int64
	N      int64
	Seed   int64
	Length int64
}

// GenerateInverseLCG Обратный линейный конгруэнтный генератор
func GenerateInverseLCG(seed, a, c, n, length int64) ([]int64, Params) {
	sequence := make([]int64, 0, length) // Создает пустой срез типа []int64 с начальной длиной 0 и емкостью length для хранения последовательности сгенерированных чисел.
	x := seed                            // Присваивает переменной x начальное значение seed начальная точка для генерации последовательности.
	mod := big.NewInt(n)                 // Создает новый объект типа *big.Int, представляющий модуль n для операций взятия остатка (модульной арифметики).
	aBig := big.NewInt(a)                // Создает новый объект типа *big.Int, представляющий множитель a участвует в расчетах обратного линейного конгруэнтного генератора.
	cBig := big.NewInt(c)                // Создает новый объект типа *big.Int, представляющий константу сдвига c добавляется к произведению в каждом шаге генерации.

	// Основной цикл для генерации последовательности псевдослучайных чисел
	for i := int64(0); i < length; i++ {
		// Создается новый объект *big.Int из текущего значения x
		xBig := big.NewInt(x)
		// Определяется обратный элемент xBig по модулю mod
		// Если обратный элемент не найден, выполнение прерывается
		inv := new(big.Int).ModInverse(xBig, mod)
		if inv == nil {
			break // Прерывание цикла, если обратный элемент не существует
		}
		xBig.Mul(aBig, inv)            // Выполняется умножение aBig на inv и результат сохраняется в xBig
		xBig.Add(xBig, cBig)           // К результату прибавляется cBig
		xBig.Mod(xBig, mod)            // Результат приводится по модулю mod
		x = xBig.Int64()               // Полученное значение преобразуется обратно в int64 и присваивается x
		sequence = append(sequence, x) // Новое значение добавляется в конец среза sequence
	}

	return sequence, Params{
		A:      a,
		C:      c,
		N:      n,
		Seed:   seed,
		Length: length,
	}
}
