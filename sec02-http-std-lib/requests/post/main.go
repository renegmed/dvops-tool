package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	resp, err := http.Post("https://httpbin.org/post",
		"text/plain",
		strings.NewReader("this is the request content"))

	if err != nil {
		log.Fatalln("Unable to get response")
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("Unable to read content")
	}

	fmt.Println(string(content))
}

/*

$ go run main.go
{
  "args": {},
  "data": "this is the request content",
  "files": {},
  "form": {},
  "headers": {
    "Accept-Encoding": "gzip",
    "Connection": "close",
    "Content-Length": "27",
    "Content-Type": "text/plain",
    "Host": "httpbin.org",
    "User-Agent": "Go-http-client/1.1"
  },
  "json": null,
  "origin": "100.12.69.210",
  "url": "https://httpbin.org/post"
}

*/
