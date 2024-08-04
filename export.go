package main

import (
	"fmt"
	"net/http"
	"strings"
)

//export Getconv
func Getconv() [3]string {
	var body string
	resp, err := http.Get("https://v6.exchangerate-api.com/v6/831604053925ff6a3be209cc/latest/RUB")

	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	for {
		bs := make([]byte, 4096)
		n, err := resp.Body.Read(bs)
		body = body + string(bs[:n])
		if n == 0 || err != nil {
			break
		}
	}

	var rtn [3]string

	var strE = body[strings.Index(body, "\"EUR"):] //обрезаем строку до "EUR"
	rtn[0] = strE[:strings.Index(strE, ",")]       //обрезаем строку после " "EUR": x.xxxxx "

	var strJ = body[strings.Index(body, "\"JPY"):] //обрезаем строку до "JPY"
	rtn[1] = strJ[:strings.Index(strJ, ",")]       //обрезаем строку после " "JPY": x.xxxxx "

	var strU = body[strings.Index(body, "\"USD"):] //обрезаем строку до "USD"
	rtn[2] = strU[:strings.Index(strU, ",")]       //обрезаем строку после " "USD": x.xxxxx "

	return rtn
}
func main() {
	fmt.Println(Getconv())
}
