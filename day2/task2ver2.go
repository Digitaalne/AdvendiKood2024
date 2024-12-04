package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		orgNumbers := strings.Split(scanner.Text(), " ")
		safe := true
		for y := 0; y < len(orgNumbers); y++ {
			safe = true
			numbers := slices.Concat(orgNumbers[:y], orgNumbers[y+1:])
			prev := 0
			prevSign := -1
			if len(numbers) > 0 {
				prev, _ = strconv.Atoi(numbers[0])
			} else {
				continue
			}
			for i := 1; i < len(numbers); i++ {
				//fmt.Println(i)
				curr, _ := strconv.Atoi(numbers[i])
				diff := curr - prev
				signBit := (diff >> 63) & 1
				diff = abs(diff)
				if isSafe(diff, signBit, prevSign) {
					prev = curr
					prevSign = signBit
				} else {
					safe = false
					break
				}
			}
			if safe {
				sum += 1
				break
			}
		}

	}
	fmt.Println(sum)
}

func abs(value int) int {
	if value < 0 {
		return -value
	}
	return value
}

func isSafe(diff int, signBit int, prevSign int) bool {
	return diff > 0 && diff < 4 && (signBit == prevSign || prevSign == -1)
}
