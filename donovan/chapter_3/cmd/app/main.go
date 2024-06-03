package main

import (
	// "go_study/donovan/chapter_3/pkg/intnums"

	"fmt"
	"math"
)

// int - знаковые, uint - беззнаковые целочисленные типы данных
// & побитовое И
// | побитовое ИЛИ
// ^ побитовое исключение ИЛИ
// &^ сброс бита (И НЕ)
// << сдвиг влево
// >> сдвиг вправо

var f float32 = 16777216 // 1 << 24

const (
	е        = 2.71828
	Avogadro = 6.02214129e23
	Planck   = 6.62606957e34
)

func main() {
	fmt.Println(f)
	fmt.Println(f == f+1) // "true"!
	fmt.Println(е)
	fmt.Println(Avogadro)
	fmt.Println(Avogadro)

	// intnums.ShowExample()

	for x := 0; x < 8; x++ {
		fmt.Printf("x = %d ex = %8.3f\n", x, math.Exp(float64(x)))
	}

	nan := math.NaN()

	fmt.Println(nan == nan, nan < nan, nan > nan)

	fmt.Println(compute())

}

func compute() (value float64, ok bool) {
	result, failed := 15.54, false

	if failed {
		return 0, false
	}
	return result, true
}
