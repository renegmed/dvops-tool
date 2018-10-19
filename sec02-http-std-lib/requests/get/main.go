package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://httpbin.org/get?search=something")
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
/*
$ go run main.go
{
  "args": {
    "search": "something"      <---- NOTE argument
  },
  "headers": {
    "Accept-Encoding": "gzip",
    "Connection": "close",
    "Host": "httpbin.org",
    "User-Agent": "Go-http-client/1.1"
  },
  "origin": "100.12.69.210",
  "url": "https://httpbin.org/get?search=something"
}

*/
