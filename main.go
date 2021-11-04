package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		println("Error:", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, resp.Body)
}
