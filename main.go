package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		println("Error:", err)
		os.Exit(1)
	}

	bsSize := 99999
	bs := make([]byte, bsSize)

	resp.Body.Read(bs)
	fmt.Println(string(bs))
}
