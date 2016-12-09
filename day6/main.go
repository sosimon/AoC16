package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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

func readInputFile(filename string) *bufio.Scanner {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Failed to open input file")
	}
	return bufio.NewScanner(file)
}

func main() {
	lineNum := 0
	data := [][]rune{}
	scanner := readInputFile("input")
	recovered := []rune{}

	for scanner.Scan() {
		message := scanner.Text()
		for i, char := range message {
			if lineNum == 0 {
				pos := []rune{char}
				data = append(data, pos)
			} else {
				data[i] = append(data[i], char)
			}
		}
		lineNum++
	}

	for _, pos := range data {
		histo := buildHistogram(pos)
		l := newLetters(histo)
		sort.Sort(l)
		recovered = append(recovered, l[len(l)-1].Character)
	}

	fmt.Println(string(recovered))
}
