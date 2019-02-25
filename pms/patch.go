package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Create("C:/Temp/result.txt")
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()

	fmt.Fprintf(file, "Hello Readers of golangcode.com")
}
