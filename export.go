package main

import (
	"C"
	"fmt"
	"net/http"
)

//export GetBody
func GetBody(currency *C.char) *C.char {

	var body string
	curr := C.GoString(currency)
	url := fmt.Sprintf("https://v6.exchangerate-api.com/v6/831604053925ff6a3be209cc/latest/%s", curr)

	resp, err := http.Get(url)
	if err != nil {
		return C.CString("Conn error")
	}
	defer resp.Body.Close()

	for {
		bs := make([]byte, 4096)
		n, err := resp.Body.Read(bs)
		body += string(bs[:n])
		if n == 0 || err != nil {
			break
		}
	}
	return C.CString(body)
}
func main() {}
