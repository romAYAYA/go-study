package main

import (
	"go_study/donovan/chapter_2/pkg/popcount"
	"go_study/donovan/chapter_2/pkg/tempconv"
	"log"

	// "go_study/donovan/chapter_2/pkg/mathalgos"

	// "flag"
	"fmt"
	"os"
	// "strings"
	// "go_study/donovan/chapter_2/pkg/declaration"
)

// var p = f()

// var n = flag.Bool("n", false, "пропуск символа новой строки")
// var sep = flag.String("s", " ", "разделитель")

// type Celsius float64
// type Fahrenheit float64

// const (
// 	absoluteZeroC Celsius = 273.15
// 	FreezingC     Celsius = 0
// 	Boioling      Celsius = 100
// )

// func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
// func FToC(f Fahrenheit) Celsius { return Celsius(f-32) * 5 / 9 }
// func (c Celsius) String() string {return fmt.Sprintf("%gC", c)}

func f() {}

var g = "g"

var cwd string

func init() {
	var err error

	cwd, err = os.Getwd()
	if err != nil {
		log.Fatalf("Error os.Getwd: %v", err)
	}
}

func main() {
	f := "f"
	fmt.Println(f) // локальная перменная затеняет функцию уровня пакета

	fmt.Println(g) // просто перменная уровня пакета

	toUpCase("privet")

	fmt.Println(tempconv.CToF(15))
	fmt.Println(tempconv.FToC(15))
	fmt.Println(popcount.PopCount(15))

	// var c Celsius
	// var f Fahrenheit

	// fmt.Println(CToF(15))
	// fmt.Println(FToC(15))
	// fmt.Println(c == 0)
	// fmt.Println(f == 0)
	// fmt.Println(f == Fahrenheit(c))

	// medals := []string{"gold", "silver", "bronze"}

	// medals[0] = "gold"
	// medals[1] = "silver"
	// medals[2] = "bronze"

	// fmt.Println(mathalgos.Gcd(15, 10))
	// fmt.Println(mathalgos.Fib(15))
	// fmt.Println(medals)

	// p := new(int)
	// fmt.Println(*p)
	// *p = 2
	// fmt.Println(*p)

	// flag.Parse()
	// fmt.Print(strings.Join(flag.Args(), *sep))
	// if !*n {
	// 	fmt.Println()
	// }

	// v := 1
	// incr(&v)
	// fmt.Println(incr(&v))

	// fmt.Println(f() == p)
	// var x, y int

	// fmt.Println(&x == &x, &x == &y, &x == nil)

	// x := 1
	// p := &x // p имеет тип int и указывает на
	// fmt.Println(*p)
	// *p = 2 // эквивалентно присваиванию x = 2
	// fmt.Println(x)

	// const freezingF, boilingF = 32.0, 212.0

	// declaration.DeclarationExample()

	// fmt.Printf("%g°F = %g°C\n", freezingF, declaration.FToC(freezingF))

	// fmt.Printf("%g°F = %g°C\n", boilingF, declaration.FToC(boilingF))
}

// func f() *int {
// 	v := 1
// 	return &v
// }

// func incr(x *int) int {
// 	*x++
// 	return *x
// }

func toUpCase(s string) {
	for _, s := range s {
		s := s + 'A' - 'a'
		fmt.Printf("%c", s)
	}
}
