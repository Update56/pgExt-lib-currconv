package main

import (
	"C"
	"net/http"
)
import "fmt"

//export GetBody
func GetBody(currency *C.char) *C.char {

	var body string
	curr := C.GoString(currency)
	url := fmt.Sprintf("https://v6.exchangerate-api.com/v6/831604053925ff6a3be209cc/latest/%s", curr)

	resp, _ := http.Get(url)
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
