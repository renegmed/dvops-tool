package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	client := http.DefaultClient
	req, err := http.NewRequest("GET", "https://httpbin.org/get", nil)
	if err != nil {
		log.Fatalln("Unable to create request")
	}
	resp, err := client.Do(req)

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
$ go build .
$ ./client
{
  "args": {},
  "headers": {
    "Accept-Encoding": "gzip",
    "Connection": "close",
    "Host": "httpbin.org",
    "User-Agent": "Go-http-client/1.1"
  },
  "origin": "100.12.69.210",
  "url": "https://httpbin.org/get"
}

*/
