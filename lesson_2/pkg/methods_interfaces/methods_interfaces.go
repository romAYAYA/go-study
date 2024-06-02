package methodsinterfaces

import "fmt"

// интерфейс
type WalkerTalker interface {
	Walk()        // контракт интерфейса
	Talk() string // контракт интерфейса
	Eater
}

type MyIFace interface{}

type Eater interface {
	Eat() string
}

// пользовательский тип данных
type Guy struct {
	Eater
}

// метод (передача по значению)
func (g *Guy) Walk() {
	fmt.Println("I walk!")
}

func (g *Guy) Talk() string {
	return "I talk!"
}

type myInt int

func (i myInt) String() string {
	return fmt.Sprintf("Целое число: %d\n", i)
}

func ShowExmpl() {
	var g Guy
	var wt WalkerTalker = &g
	wt.Walk()
	println(wt.Talk())
	// println(wt.Eat())
	var mi MyIFace
	i := 10
	s := "str"
	mi = i
	mi = s
	_ = mi
	// g.Eat()

	var ex interface{}
	ex = "20"
	ex = "abc"
	ex = []int{1, 2, 3}
	fmt.Println(ex)

	var mn myInt = 10
	fmt.Println(mn)
}
