package structexmpl

import "fmt"

type point struct {
	lat float64
	lon float64
}

type object struct {
	name  string
	point // встраивание, обращение по имени типа
}

func ShowExmpl() {
	// o := object{}
	// o.name = "Object"
	// o.lat = 10
	// o.lon = 20
	// o.PrintPoint()

	p := point{lat: 10, lon: 20}
	fmt.Printf("%+v\n", p)
	Change(p)
	fmt.Printf("%+v\n", p)
}

// метод - функция со специальным аргументом - приемником (receiver)
// получатель указывается в скобках перед именем функции. может быть как указателем, так и значением
func (p point) PrintPoint() {
	fmt.Printf("lat:lon - %f, %f\n", p.lat, p.lon)
}

func Change(p point) {
	p.lat = 100
	p.lon = 200
}
