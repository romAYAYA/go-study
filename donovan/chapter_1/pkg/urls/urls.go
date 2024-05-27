package urls

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func SendReq() {
	// создал штуку для чтения ввода из баша
	for _, url := range os.Args[1:] {
		p := "https://"
		if !strings.HasPrefix(url, "https") {
			url = p + url
		}
		// вот так вот сделал запрос
		res, err := http.Get(url)
		// никаких трай кетчев, поэтому ошибки так обрабатываются
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			// закрываем открытые файлы, чтобы утечки не произошло
			os.Exit(1)
		}

		fmt.Println(res.StatusCode)

		// тут просто файлик создаем
		file, err := os.Create("output.txt")

		// снова обработка ошибок
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
		}
		// записываем ответ в файлик, что создали
		_, err = io.Copy(file, res.Body)
		// закрыли тело ответа
		res.Body.Close()
		// любимая обработка ошибок
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: чтение %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}
