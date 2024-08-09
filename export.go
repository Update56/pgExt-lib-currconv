package main

import "C"

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"unsafe"
)

//export Getconv
func Getconv(rates *C.float) {
	var body string
	resp, _ := http.Get("https://v6.exchangerate-api.com/v6/831604053925ff6a3be209cc/latest/RUB")
	defer resp.Body.Close()
	for {
		bs := make([]byte, 4096)
		n, err := resp.Body.Read(bs)
		body += string(bs[:n])
		if n == 0 || err != nil {
			break
		}
	}

	floats := (*[3]C.float)(unsafe.Pointer(rates))
	floats[0] = C.float(getFloat(selectCurr("\"EUR", body)))
	floats[1] = C.float(getFloat(selectCurr("\"JPY", body)))
	floats[2] = C.float(getFloat(selectCurr("\"USD", body)))
}

func getFloat(str string) float32 {
	value, err := strconv.ParseFloat(str[6:], 32)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return float32(value)
}

func selectCurr(curr string, body string) string {
	var strE = body[strings.Index(body, curr):]
	return strE[:strings.Index(strE, ",")]
}

func main() {}
