package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
)

type Letter struct {
	Character rune
	Frequency int
}

type Letters []*Letter

func (l Letters) Len() int { return len(l) }

func (l Letters) Swap(i, j int) { l[i], l[j] = l[j], l[i] }

func (l Letters) Less(i, j int) bool {
	if l[i].Frequency != l[j].Frequency {
		return l[i].Frequency > l[j].Frequency
	}
	return l[i].Character < l[j].Character
}

func newLetters(histogram map[rune]int) Letters {
	l := Letters{}
	for r, f := range histogram {
		l = append(l, &Letter{
			Character: r,
			Frequency: f,
		})
	}
	return l
}

func readInputFile(filename string) *bufio.Scanner {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal("Failed to open input file")
	}
	return bufio.NewScanner(file)
}

func checksum(encrypted string) string {
	re := regexp.MustCompile(`\[([a-z]+)\]`)
	matches := re.FindAllStringSubmatch(encrypted, -1)
	return matches[0][1]
}

func sectorID(encrypted string) int {
	re := regexp.MustCompile(`(\d+)`)
	matches := re.FindAllStringSubmatch(encrypted, -1)
	id, err := strconv.Atoi(matches[0][1])
	if err != nil {
		log.Fatalf("Failed to convert %s to int", matches[0][1])
	}
	return id
}

func buildHistogram(letters []rune) map[rune]int {
	histogram := make(map[rune]int)
	for _, r := range letters {
		if _, ok := histogram[r]; ok {
			histogram[r]++
		} else {
			histogram[r] = 1
		}
	}
	return histogram
}

func calcChecksum(encrypted string) string {
	re := regexp.MustCompile(`([a-z]+)-`)
	matches := re.FindAllStringSubmatch(encrypted, -1)
	letters := []rune{}
	// combine matches into one string
	for _, m := range matches {
		for _, letter := range m[1] {
			letters = append(letters, letter)
		}
	}
	// create histogram map with letter as key and frequency as value
	histogram := buildHistogram(letters)
	l := newLetters(histogram)
	// sort by letter frequency (alphebetical order as tie-breaker)
	sort.Sort(l)
	checksum := []rune{}
	// top 5 is the checksum
	for i := 0; i < 5; i++ {
		checksum = append(checksum, l[i].Character)
	}
	return string(checksum)
}

func decrypt(encryptedname string) string {
	sectorId := sectorID(encryptedname)
	re := regexp.MustCompile(`([a-z]+)-`)
	matches := re.FindAllStringSubmatch(encryptedname, -1)
	letters := []rune{}
	// combine matches into one string
	for _, m := range matches {
		for _, letter := range m[1] {
			letters = append(letters, letter)
		}
	}
	decrypted := []rune{}
	for _, l := range letters {
		decrypted = append(decrypted, shift(l, sectorId))
	}
	return string(decrypted)
}

func shift(letter rune, sectorId int) rune {
	codepoint := letter + rune(sectorId%26)
	// a-z == 97-122
	if codepoint > 122 {
		codepoint -= 26
	}
	return rune(codepoint)
}

func main() {
	scanner := readInputFile("input")
	sectorIDSum := 0
	for scanner.Scan() {
		encryptedName := scanner.Text()
		//fmt.Printf("Calculated checksum: %s\n", calcChecksum(encryptedName))
		//fmt.Printf("Provided checksum: %s\n", checksum(encryptedName))
		//fmt.Printf("Secter ID: %d\n", sectorID(encryptedName))
		if calcChecksum(encryptedName) == checksum(encryptedName) {
			sectorIDSum += sectorID(encryptedName)
			fmt.Println(encryptedName)
			fmt.Printf("Decrypted name: %s\n", decrypt(encryptedName))
		}

	}
	fmt.Println(sectorIDSum)
}
