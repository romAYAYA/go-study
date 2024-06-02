package exmpls

import (
	"fmt"
	"reflect"
)

func Range() {
	slice := []int{10, 20, 30}
	arr := [...]int{10, 20, 30}
	str := "ABC"

	for i, v := range slice {
		fmt.Printf("%d %d\n", i, v)
	}
	fmt.Println("\n\n")

	for i, v := range arr {
		fmt.Printf("%d %d\n", i, v)
	}
	fmt.Println("\n\n")

	for i, v := range str {
		fmt.Printf("%d %d, %s\n", i, v, string(v))
	}

}

func Vars() {
	var i int
	var j int32 = 10
	k := 20
	fmt.Println(i, j, k)
}

func Types() {
	type circle struct {
		x, y int
		r    int
	}

	c := circle{
		x: 10,
		y: 20,
		r: 30,
	}
	fmt.Printf("%+v\n", c)

	type myString string
	var s string = "ABC"
	var s1 myString = myString(s)
	fmt.Printf("Type 1: %v\tType 2: %v\t\n", reflect.TypeOf(s), reflect.TypeOf(s1))
}

func Pointers() {
	var s string = "ABC"
	var pointer *string = &s
	fmt.Println(*pointer, pointer)
	fmt.Printf("Value: %v\tPointer: %v\tRef: %v\t\n", s, pointer, &s)
}

func Scopes() {
	x := 1
	{
		x := 2
		fmt.Printf("x = %v\n", x)
	}
	fmt.Printf("x = %v\n", x)
}

func Index(s string, b byte) int {
	for i := range s {
		if s[i] == b {
			return i
		}
	}
	return -1
}

func PointersAdv() {
	var i int = 10 // переменная
	var p *int     // указатель
	p = &i         // взятие адреса. указателю присваивается ссылка
	var p1 *interface{}
	_ = p1
	fmt.Println(i)
	i++
	fmt.Println(i)
	*p++ // разыменовывание указателя
	fmt.Println(i)
	*(&i)++
	fmt.Println(i)
}

func PntrRefValExmpl() {
	var p *int 		// указатель на переменную типа int
	var i int = 20	// переменная со значением 20
	addr := &i		// ссылка на переменную
	p = addr		// корректно 
	fmt.Println(*p)	// разыменование указателя (получение значения)
}
