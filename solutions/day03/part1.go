package day03

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func Part1() (string, error) {
	file, err := os.OpenFile("solutions/day03/input.txt", os.O_RDONLY, 0600)
	if err != nil {
		return "", err
	}

	answer, puzzleErr := solvePart1(bufio.NewReader(file))
	return fmt.Sprintf("%d", answer), puzzleErr
}

func solvePart1(file *bufio.Reader) (uint32, error) {
	memoryScanner := bufio.NewScanner(file)
	total := 0
	content := ""
	for memoryScanner.Scan() {
		content += memoryScanner.Text()
	}

	re := regexp.MustCompile(`mul\(\d+,\d+\)`)
	matches := re.FindAll([]byte(content), -1)

	for _, match := range matches {
		numbers := regexp.MustCompile(`\d+`)
		numberMatches := numbers.FindAll(match, 2)

		first, _ := strconv.Atoi(string(numberMatches[0]))
		second, _ := strconv.Atoi(string(numberMatches[1]))
		total += first * second
	}

	return uint32(total), nil
}
