package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func strToInt(str []string) []int {
	out := []int{}
	for _, s := range str {
		if len(s) == 0 {
			continue
		}
		num, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalf("Error converting %s to int", s)
		}
		out = append(out, num)
	}
	return out
}

func main() {
	var possibleCount = 0
	file, err := os.Open("input")
	if err != nil {
		log.Fatal("Failed to open input file")
	}
	scanner := bufio.NewScanner(file)
	count := 0
	var lines [3][]int
	triangles := [][]int{}
	for scanner.Scan() {
		raw := scanner.Text()
		line := strings.Split(raw, " ")
		lengths := strToInt(line)
		//sort.Ints(lengths)
		lines[count%3] = lengths
		if (count+1)%3 == 0 {
			for i := 0; i < 3; i++ {
				triangle := []int{}
				for j := 0; j < 3; j++ {
					triangle = append(triangle, lines[j][i])
				}
				triangles = append(triangles, triangle)
				fmt.Println(triangle)
			}
		}
		count++

	}
	for _, t := range triangles {
		sort.Ints(t)
		if t[0]+t[1] > t[2] {
			possibleCount++
		}
	}
	fmt.Printf("Total number of possible triangles: %d\n", possibleCount)
}
