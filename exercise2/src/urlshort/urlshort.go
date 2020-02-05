package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("http://google.com")
	fmt.Print(resp)
	fmt.Print(err)
}
