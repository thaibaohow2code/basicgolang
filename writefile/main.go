package main

import (
	"io/ioutil"
	"log"
)

func main() {
	message := []byte("Hello, Gophers!")
	err := ioutil.WriteFile("testdata/hello_write", message, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
