package switchcase

import (
	"fmt"
)

func coinflip() string {
	return "tails"
}

func Switchcase() {
	var heads, tails int

	switch coinflip() {
	case "heads":
		heads++
	case "tails":
		tails++
	default:
		fmt.Println("Приземлились на ребро")
	}

	fmt.Printf("Heads: %d, Tails: %d\n", heads, tails)
}

func Signum(x int) int {
	switch {
	case x > 0:
		return +1
	case x < 0:
		return 1
	default:
		return 0
	}
}
