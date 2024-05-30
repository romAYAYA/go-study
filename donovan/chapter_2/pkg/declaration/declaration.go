package declaration

import "fmt"

const boilingF = 212.0

func DeclarationExample() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("Температура кипения = %g°F или %g°C\n", f, c)
}

func FToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
