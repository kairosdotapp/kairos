package main

import (
	"bufio"
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

	scanner := bufio.NewScanner(f)

	p := newParser(scanner)

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
