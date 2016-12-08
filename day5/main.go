package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readInputFile(filename string) *bufio.Scanner {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal("Failed to open input file")
	}
	return bufio.NewScanner(file)
}

func main() {
	scanner := readInputFile("input")
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
