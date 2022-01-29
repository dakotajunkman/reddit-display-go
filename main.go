package main

import (
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
)

func main() {
	fmt.Println("This works")
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
}