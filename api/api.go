package api

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	methods := r.URL.Query()
	method := methods.Get("method")
	num1, _ := strconv.Atoi(methods.Get("a"))
	num2, _ := strconv.Atoi(methods.Get("b"))

	switch method {
	case "mul":
		fmt.Fprintln(w, "Nhan hai so")
		fmt.Fprintln(w, num1*num2)
	case "div":
		fmt.Fprintln(w, "Chia hai so")
		if num1 < num2 {
			fmt.Fprintln(w, float64(num2)/float64(num1))
		} else {
			fmt.Fprintln(w, float64(num1)/float64(num2))
		}
	case "sum":
		fmt.Fprintln(w, "Cong hai so")
		fmt.Fprintln(w, num1+num2)
	case "sub":
		fmt.Fprintln(w, "Tru hai so")
		if num1 < num2 {
			fmt.Fprintln(w, num2-num1)
		} else {
			fmt.Fprintln(w, num1-num2)
		}
	default:
		fmt.Fprintln(w, "Khong hop le")
	}
}

func HandleRequest() {
	http.HandleFunc("/", HomePage)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
