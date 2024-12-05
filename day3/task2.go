package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	str := string(file)
	delRegex, _ := regexp.Compile("don't\\(\\)((?:.|\\s)*?|)(?:do\\(\\))")
	r, _ := regexp.Compile("mul\\((\\d{1,3}),(\\d{1,3})\\)")

	sum := 0
	str = delRegex.ReplaceAllString(str, "")
	matches := r.FindAllStringSubmatch(str, -1)
	for i := 0; i < len(matches); i++ {
		first, _ := strconv.Atoi(matches[i][1])
		second, _ := strconv.Atoi(matches[i][2])
		sum += first * second
	}
	fmt.Println(sum)
}
