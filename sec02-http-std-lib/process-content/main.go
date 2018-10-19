package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type GetResponse struct {
	Origin  string            `json:"origin"`
	URL     string            `json:"url"`
	Headers map[string]string `json:"headers"`
}

func (r *GetResponse) ToString() string {
	s := fmt.Sprintf("GET Response\nOrigin IP: %s\nRequest URL: %s\n",
		r.Origin, r.URL)
	for k, v := range r.Headers {
		s += fmt.Sprintf("Header: %s = %s\n", k, v)
	}
	return s
}
func main() {
	resp, err := http.Get("https://httpbin.org/get")
	if err != nil {
		log.Fatalln("Unable to read response")
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("Unable to read content")
	}
	respContent := GetResponse{}
	json.Unmarshal(content, &respContent)

	fmt.Println(respContent.ToString())
}

/*

$ go run main.go
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

GET ResponseOrigin IP: 100.12.69.210
Request URL: https://httpbin.org/get
Header: Accept-Encoding = gzip
Header: Connection = close
Header: Host = httpbin.org
Header: User-Agent = Go-http-client/1.1

*/
