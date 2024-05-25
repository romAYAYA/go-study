package main

import (
	"fmt"
	"os"
	"strconv"
	// "strings"
)

func main() {
	var s, sep string

	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i] + " " + strconv.Itoa(i)
		sep = "\n"
	}

	fmt.Println(s)
}

// func main() {
// 	var s, sep string

// 	for _, arg := range os.Args[1:] {
// 		s += sep + arg
// 		sep = "\n"
// 	}

// 	fmt.Println(os.Args[0]+"\n", s)
// }

// func main() {
// 	fmt.Println(strings.Join(os.Args[1:], " "))
// }
