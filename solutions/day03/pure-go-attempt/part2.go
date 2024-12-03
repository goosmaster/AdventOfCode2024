package day03

import (
	"bufio"
	"fmt"
	"os"
)

func Part2() (string, error) {
	file, err := os.OpenFile("solutions/day03/input.txt", os.O_RDONLY, 0600)
	if err != nil {
		return "", err
	}

	answer, puzzleErr := solvePart1(bufio.NewReader(file))
	return fmt.Sprintf("%d", answer), puzzleErr
}

func solvePart2(file *bufio.Reader) (uint32, error) {
	memoryScanner := bufio.NewScanner(file)
	var instructions [][]int
	memoryBuffer := make([]byte, 3)
	byteChunk := make([]byte, 8)
	isMultiply := false

	for memoryScanner.Scan() {
		line := memoryScanner.Bytes()
		for _, memory := range line {
			if isMultiply == false {
				isMultiply = isMultiplyInstruction(memory, &memoryBuffer)
			}
			if isMultiply == false {
				continue
			}
			if len(byteChunk) == 0 {
				byteChunk = memoryBuffer // mul(
				continue
			}
			if len(byteChunk) < 12 && memory != ')' {
				byteChunk = append(byteChunk, memory)
				continue
			}
			if memory == ')' {
				byteChunk = append(byteChunk, memory)
			}
			if len(byteChunk) == 12 || memory == ')' {
				// mul(xxx,xxx)
				numbers, err := parseMultiplyNumbers(uint8(3), ',', ')', byteChunk)
				if err == nil {
					instructions = append(instructions, numbers)
				} else {
					fmt.Printf("%s\n", err.Error())
				}
			}
			byteChunk = []byte{}
			memoryBuffer = []byte{}
			isMultiply = false
		}
	}

	total := uint32(0)
	for _, instruction := range instructions {
		total += uint32(instruction[0] * instruction[1])
	}
	return total, nil
}
