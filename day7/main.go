package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

// THIS DOES NOT WORK
// Golang regexp does not support back references
// Switched to Python for this day's puzzle
// See day7.py
func supportsTLS(ip string) {
	abbaPattern := regexp.MustCompile(`(\w)(\w)\2\1`)
	matches := abbaPattern.FindAllStringSubmatch(ip, -1)
	fmt.Println(matches)
}

func readInputFile(filename string) *bufio.Scanner {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Failed to open input file")
	}
	return bufio.NewScanner(file)
}

func main() {
	scanner := readInputFile("input")
	for scanner.Scan() {
		ip := scanner.Text()
		fmt.Println(ip)
		supportsTLS(ip)
	}
}
