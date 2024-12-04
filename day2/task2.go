package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input_test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		safeCount := 0
		retry := false
		numbers := strings.Split(scanner.Text(), " ")
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
				safeCount += 1
				if !retry && i < 5 && safeCount == 2 {
					retry = true
					safeCount = 1
					i = 1
					prev, _ = strconv.Atoi(numbers[1])
					prevSign = -1
					/* next, _ := strconv.Atoi(numbers[i+1])
					altDiff := next - curr
					altSignBit := (altDiff >> 63) & 1
					if isSafe(altDiff, altSignBit, -1) {
						prev = next
						prevSign = altSignBit
						i += 1
					} */
				}
				if safeCount > 2 {
					break
				}
			}
		}
		//fmt.Println(scanner.Text(), "=", safeCount)
		if safeCount > 1 {
			fmt.Println(scanner.Text())
		}
		if safeCount < 2 {
			sum += 1
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
