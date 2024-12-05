package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	r, _ := regexp.Compile("mul\\((\\d{1,3}),(\\d{1,3})\\)")

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		matches := r.FindAllStringSubmatch(scanner.Text(), -1)
		for i := 0; i < len(matches); i++ {
			first, _ := strconv.Atoi(matches[i][1])
			second, _ := strconv.Atoi(matches[i][2])
			sum += first * second
		}
	}
	fmt.Println(sum)
}
