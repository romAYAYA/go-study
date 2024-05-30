package main

import (
	"flag"
	"fmt"
	"strings"
	// "go_study/donovan/chapter_2/pkg/declaration"
)

// var p = f()

var n = flag.Bool("n", false, "пропуск символа новой строки")
var sep = flag.String("s", " ", "разделитель")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}

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
