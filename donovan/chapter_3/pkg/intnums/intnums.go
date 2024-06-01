package intnums

import "fmt"

// x<<n эквивалетно умножению на 2**n
// x>>n эквивалетно делению на 2**n

var x uint8 = 1<<1 | 1<<5
var y uint8 = 1<<1 | 1<<2

func ShowExample() {
	fmt.Printf("%08b\n", x)    // 00100010 множество {1, 5}
	fmt.Printf("%08b\n", y)    // 00000110 множество {1, 2}
	fmt.Printf("%08b\n", x|y)  // 00100110 пересечение {1}
	fmt.Printf("%08b\n", x&y)  // 00000010 объединение {1, 2, 5}
	fmt.Printf("%08b\n", x^y)  // 00100100 симметричная разность {2, 5}
	fmt.Printf("%08b\n", x&^y) // 00100000 разность {5}

	for i := uint(0); i < 8; i++ {
		if x&(1<<i) != 0 { // проверка принадлежности множеству
			fmt.Println(i) // 1, 5
		}
	}

	fmt.Printf("%08b\n", x<<1) // 01000100, множетсво {2, 6}
	fmt.Printf("%08b\n", x>>1) // 00010001, множетсво {0, 4}

	o := 0666
	fmt.Printf("%d %[1]o %#[1]o\n", o)

	x := int64(0xdeadbeef)
	fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", x)

	ascii := 'а'
	Unicode := '👽'
	newline := '\n'
	fmt.Printf("%d %[1]c %[1]q\n", ascii)
	fmt.Printf("%d %[1]c %[1]q\n", Unicode)
	fmt.Printf("%d %[1]q\n", newline)
}
