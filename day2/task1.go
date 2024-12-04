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
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	safe := true
	for scanner.Scan() {
		safe = true
		numbers := strings.Split(scanner.Text(), " ")
		prev := 0
		prevSign := -1
		if len(numbers) > 0 {
			prev, _ = strconv.Atoi(numbers[0])
		} else {
			continue
		}
		for i := 1; i < len(numbers); i++ {
			curr, _ := strconv.Atoi(numbers[i])
			diff := curr - prev
			signBit := (diff >> 63) & 1
			diff = abs(diff)
			if diff > 0 && diff < 4 && (signBit == prevSign || prevSign == -1) {
				prev = curr
				prevSign = signBit
			} else {
				safe = false
				break
			}
		}
		if safe {
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
