package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("invoice-cli")

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal("Error opening file: ", err)
	}

	defer f.Close()

	p := newTimedotParser(f)

	for {
		e, err := p.scan()

		if err != nil {
			log.Fatal(err)
		}

		if e == nil {
			break
		}

		fmt.Println(e)
	}
}
