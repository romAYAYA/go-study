package main

import (
	"go_study/lesson_1/demo_web/pkg/stringutils"
	"net/http"
)

func main() {
	http.ListenAndServe(
		":8080",
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			s := r.URL.Query()["s"]
			rev := stringutils.RevStr(s[0])
			w.Write([]byte(rev))
		},
		),
	)
}
