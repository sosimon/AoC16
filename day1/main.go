package main

import "fmt"
import "strconv"
import "math"

var directions = []string{"N", "E", "S", "W"}

var input = []string{"R4", "R4", "L1", "R3", "L5", "R2", "R5", "R1", "L4", "R3", "L5", "R2", "L3", "L4", "L3", "R1", "R5", "R1", "L3", "L1", "R3", "L1", "R2", "R2", "L2", "R5", "L3", "L4", "R4", "R4", "R2", "L4", "L1", "R5", "L1", "L4", "R4", "L1", "R1", "L2", "R5", "L2", "L3", "R2", "R1", "L194", "R2", "L4", "R49", "R1", "R3", "L5", "L4", "L1", "R4", "R2", "R1", "L5", "R3", "L5", "L4", "R4", "R4", "L2", "L3", "R78", "L5", "R4", "R191", "R4", "R3", "R1", "L2", "R1", "R3", "L1", "R3", "R4", "R2", "L2", "R1", "R4", "L5", "R2", "L2", "L4", "L2", "R1", "R2", "L3", "R5", "R2", "L3", "L3", "R3", "L1", "L1", "R5", "L4", "L4", "L2", "R5", "R1", "R4", "L3", "L5", "L4", "R5", "L4", "R5", "R4", "L3", "L2", "L5", "R4", "R3", "L3", "R1", "L5", "R5", "R1", "L3", "R2", "L5", "R5", "L3", "R1", "R4", "L5", "R4", "R2", "R3", "L4", "L5", "R3", "R4", "L5", "L5", "R4", "L4", "L4", "R1", "R5", "R3", "L1", "L4", "L3", "L4", "R1", "L5", "L1", "R2", "R2", "R4", "R4", "L5", "R4", "R1", "L1", "L1", "L3", "L5", "L2", "R4", "L3", "L5", "L4", "L1", "R3"}

//var input = []string{"R5", "L5", "R5", "R3"}

func main() {
	var curDirectionIdx = 0
	var x = 0
	var y = 0
	var history = map[string]int{
		"0,0": 1,
	}
	for _, s := range input {
		dir := string(s[0])
		dist, _ := strconv.Atoi(string(s[1:]))
		fmt.Printf("%s: %d\n", dir, dist)
		if dir == "R" {
			curDirectionIdx++
		} else if dir == "L" {
			curDirectionIdx--
		}
		curDirectionIdx = (curDirectionIdx + 4) % 4
		curDirection := directions[curDirectionIdx]
		targetX := x
		targetY := y
		switch curDirection {
		case "N":
			targetY += dist
		case "S":
			targetY -= dist
		case "E":
			targetX += dist
		case "W":
			targetX -= dist
		}
		for {
			if x == targetX && y == targetY {
				break
			}
			switch curDirection {
			case "N":
				y++
			case "S":
				y--
			case "E":
				x++
			case "W":
				x--
			}
			fmt.Printf("Location: x = %d, y = %d\n", x, y)
			key := fmt.Sprintf("%d,%d", x, y)
			if _, ok := history[key]; ok {
				fmt.Println("========== !BEEN HERE BEFORE! ==========")
				history[key]++
			} else {
				history[key] = 1
			}
		}
		fmt.Printf("Location: x = %d, y = %d\n", x, y)
		d := math.Abs(float64(x)) + math.Abs(float64(y))
		fmt.Printf("Distance to HQ: %f\n", d)
	}
}
