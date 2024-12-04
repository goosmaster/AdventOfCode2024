package day04

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func Part2() (string, error) {
	file, err := os.OpenFile("solutions/day03/input.txt", os.O_RDONLY, 0600)
	if err != nil {
		return "", err
	}

	answer, puzzleErr := solvePart2(bufio.NewReader(file))
	return fmt.Sprintf("%d", answer), puzzleErr
}

func solvePart2(file *bufio.Reader) (uint32, error) {
	memoryScanner := bufio.NewScanner(file)
	total := 0
	content := ""
	for memoryScanner.Scan() {
		content += memoryScanner.Text()
	}

	re := regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`)
	matches := re.FindAll([]byte(content), -1)

	flag := true
	for _, match := range matches {

		switch string(match) {
		case "do()":
			flag = true
		case "don't()":
			flag = false
		default:
			if flag {
				numbers := regexp.MustCompile(`\d+`)
				numberMatches := numbers.FindAll(match, 2)

				first, _ := strconv.Atoi(string(numberMatches[0]))
				second, _ := strconv.Atoi(string(numberMatches[1]))
				total += first * second
			}
		}
	}

	return uint32(total), nil
}
