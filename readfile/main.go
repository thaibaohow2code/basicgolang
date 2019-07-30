package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	content, err := ioutil.ReadFile("./testdata/hello")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("File contents", string(content))

}
