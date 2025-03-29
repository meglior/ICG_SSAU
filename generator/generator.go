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
	sequence := make([]int64, 0, length)
	x := seed
	mod := big.NewInt(n)
	aBig := big.NewInt(a)
	cBig := big.NewInt(c)

	for i := int64(0); i < length; i++ {
		xBig := big.NewInt(x)
		inv := new(big.Int).ModInverse(xBig, mod)
		if inv == nil {
			break // Обратный элемент не существует
		}
		xBig.Mul(aBig, inv)
		xBig.Add(xBig, cBig)
		xBig.Mod(xBig, mod)
		x = xBig.Int64()
		sequence = append(sequence, x)
	}

	return sequence, Params{
		A:      a,
		C:      c,
		N:      n,
		Seed:   seed,
		Length: length,
	}
}
