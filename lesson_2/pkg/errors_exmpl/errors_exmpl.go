package errorsexmpl

import (
	"fmt"
	"log"
	"os"
)

func ShowExmpl() {
	val, err := envVar("несуществующее_имя_переменной")
	if err != nil {
		log.Fatal(err)
	}
	_ = val
}

func envVar(name string) (string, error) {
	val := os.Getenv(name)
	if val == "" {
		return val, fmt.Errorf("переменная с имененем %s не найдена", name)
	}
	return val, nil
}
