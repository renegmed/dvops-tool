package main

import (
	nap "devops_tools/sec03-nap/cmd"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

/*

	$ go install ./cmd/...
	$ nap -h
	$ nap -list

*/
var api = nap.NewAPI("https://httpbin.org")

func main() {
	list := flag.Bool("list", false, "Get list of all API resources")
	flag.Parse()

	if *list { // if call is $ nap -list
		fmt.Println("Available Resources:")
		for _, name := range api.ResourceNames() {
			fmt.Println(name)
		}
		return
	}
	/*
				For example
				$ nap get
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
	resource := os.Args[1]
	if err := api.Call(resource, nil); err != nil {
		log.Fatalln(err)
	}

}

func init() {
	router := nap.NewRouter()
	router.RegisterFunc(200, func(resp *http.Response, _ interface{}) error {
		defer resp.Body.Close()
		content, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		fmt.Println(string(content))
		return nil
	})
	api.AddResource("get", nap.NewResource("/get", "GET", router))
}
