package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
)

var input = "uqwqemis"
var testInput = "abc"
var index = 0

//var charCount = 0

func getNextString(input string) string {
	index++
	return input + strconv.Itoa(index)
}

func hasUnknown(password [8]string) bool {
	for _, char := range password {
		if char == "" {
			return true
		}
	}
	return false
}

func main() {
	password := [8]string{}
	//for charCount < 8 {
	for hasUnknown(password) {
		toBeHashed := getNextString(input)
		hash := md5.Sum([]byte(toBeHashed))
		hashStr := fmt.Sprintf("%x", hash)
		if hashStr[:5] == "00000" {
			//charCount++
			fmt.Printf("%s - %s - %s - %s\n", toBeHashed, hashStr, hashStr[5:6], hashStr[6:7])
			pos, err := strconv.Atoi(hashStr[5:6])
			if err != nil {
				continue
			}
			if (pos >= 0 && pos < 8) && (password[pos] == "") {
				password[pos] = hashStr[6:7]
				fmt.Printf("  Password: %q\n", password)
			}
		}
	}
}
