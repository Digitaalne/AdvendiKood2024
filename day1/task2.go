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
	fmt.Println("Hello, World!")
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var firstColumn []int
	secondColumnMap := make(map[int]int)
	for scanner.Scan() {
		var line = scanner.Text()
		var values = strings.Split(line, " ")
		first, _ := strconv.Atoi(values[0])
		second, _ := strconv.Atoi(values[len(values)-1])
		firstColumn = append(firstColumn, first)
		secondColumnMap[second] = secondColumnMap[second] + 1
	}
	sum := 0
	for i := 0; i < len(firstColumn); i++ {
		cur := firstColumn[i]
		sum += secondColumnMap[cur] * cur
	}
	fmt.Println(sum)
}
