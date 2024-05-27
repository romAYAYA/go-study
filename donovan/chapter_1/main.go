package main

import (
	"go_study/donovan/chapter_1/pkg/urls"
	// "bufio"
	// "fmt"
	// "image"
	// "image/color"
	// "image/gif"
	// "io"
	// "math"
	// "math/rand"
	// "os"

	// "strings"
	// "strconv"
	// "strings"
)

func main() {
	urls.SendReq()
	// var s, sep string

	// for i := 0; i < len(os.Args); i++ {
	// 	s += sep + os.Args[i] + " " + strconv.Itoa(i)
	// 	sep = "\n"
	// }

	// fmt.Println(s)

	// countRepStrs()
	// dup()
	// dup2()

	// file, err := os.Create("lissajous.gif")
    // if err != nil {
    //     panic(err)
    // }
    // defer file.Close()

    // Вызываем функцию lissajous для создания анимации
    // lissajous(file)
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

// func countRepStrs() {
// 	counts := make(map[string]int)
// 	input := bufio.NewScanner(os.Stdin)

// 	for input.Scan() {
// 		counts[input.Text()]++
// 	}

// 	for line, n := range counts {
// 		if n > 1 {
// 			fmt.Printf("%d\t%s\n", n, line)
// 		}
// 	}
// }

// func dup() {
// 	counts := make(map[string]int)
// 	files := os.Args[1:]

// 	if len(files) == 0 {
// 		countLines(os.Stdin, counts)
// 	} else {
// 		for _, arg := range files {
// 			f, err := os.Open(arg)
// 			if err != nil {
// 				fmt.Fprintf(os.Stderr, "dup: %v\n", err)
// 				continue
// 			}
// 			countLines(f, counts)
// 			f.Close()
// 		}
// 	}
// 	for line, n := range counts {
// 		if n > 1 {
// 			fmt.Printf("%d\t%s\t%v\n", n, line, os.Args[1])
// 		}
// 	}
// }

// func countLines(f *os.File, counts map[string]int) {
// 	input := bufio.NewScanner(f)
// 	for input.Scan() {
// 		counts[input.Text()]++
// 	}
// }

// func dup2() {
// 	counts := make(map[string]int)
// 	for _, filename := range os.Args[1:] {
// 		data, err := os.ReadFile(filename)
// 		if err != nil {
// 			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
// 			continue
// 		}
// 		for _, line := range strings.Split(string(data), "\n") {
// 			counts[line]++
// 		}
// 	}
// 	for line, n := range counts {
// 		fmt.Printf("%d\t%s\n", n, line)
// 	}
// }

// var pallete = []color.Color{color.White, color.Black}

// const (
// 	whiteIndex = 0
// 	blackIndex = 0
// )

// func lissajous(out io.Writer) {
// 	const (
// 		cycles = 5 // кол-во полных колебаний x
// 		res = 0.001 // угловое разрешение
// 		size = 100 // канва изображения охватывает [size...+size]
// 		nframes = 64 // количество кадров анимации
// 		delay = 8 // задержка между кадрами (единица - 10 мс)
// 	)
// 	freq := rand.Float64() * 3.0
// 	anim := gif.GIF{LoopCount: nframes}
// 	phase := 0.0
// 	for i := 0; i < nframes; i++ {
// 		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
// 		img := image.NewPaletted(rect, pallete)
// 		for t := 0.0; t < cycles * 2*math.Pi; t+=res {
// 			x := math.Sin(t)
// 			y :=math.Sin(t*freq+phase)
// 			img.SetColorIndex(size + int(x*size+0.5), size +int(y*size+0.5), blackIndex)
// 		}
// 		phase += 0.1
// 		anim.Delay = append(anim.Delay, delay)
// 		anim.Image = append(anim.Image, img)
// 	}
// 	gif.EncodeAll(out, &anim)
// }