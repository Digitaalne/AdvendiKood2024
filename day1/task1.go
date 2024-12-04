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

func main() {
	fmt.Println("Hello, World!")
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var firstColumn []int
	var secondColumnt []int
	for scanner.Scan() {
		var line = scanner.Text()
		var values = strings.Split(line, " ")
		first, _ := strconv.Atoi(values[0])
		second, _ := strconv.Atoi(values[len(values)-1])
		firstColumn = append(firstColumn, first)
		secondColumnt = append(secondColumnt, second)
	}

	sort.Ints(firstColumn)
	sort.Ints(secondColumnt)
	sum := 0
	for i := 0; i < len(firstColumn); i++ {
		fmt.Println(firstColumn[i], secondColumnt[i])
		if firstColumn[i] > secondColumnt[i] {
			sum += firstColumn[i] - secondColumnt[i]
		} else {
			sum += secondColumnt[i] - firstColumn[i]
		}
	}

	fmt.Println(sum)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
