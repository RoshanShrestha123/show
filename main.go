package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	args := os.Args

	if len(args[1:]) == 0 {
		log.Fatal("Must provide the file name")

	}

	fileName := args[1:]

	data, err := os.ReadFile(fileName[0])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))
}
