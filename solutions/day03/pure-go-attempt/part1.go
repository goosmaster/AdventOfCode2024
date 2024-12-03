package day03

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func Part1() (string, error) {
	file, err := os.OpenFile("solutions/day03/input.txt", os.O_RDONLY, 0600)
	if err != nil {
		return "", err
	}

	answer, puzzleErr := solvePart1(bufio.NewReader(file))
	return fmt.Sprintf("%d", answer), puzzleErr
}

// 10617205 == too low
// 172796789 == too low
// missing:
//
//		mul(949,117)
//		mul(589,854)
//		mul(496,455)
//		mul(367,970)
//	 	mul(147,254)
//		mul(119,163)
//		mul(112,401)
//		mul(502,27)
//		mul(294,14)
func solvePart1(file *bufio.Reader) (uint32, error) {
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

func isMultiplyInstruction(newByte byte, lastThreeBytes *[]byte) bool {
	*lastThreeBytes = append(*lastThreeBytes, newByte)
	if bytes.Equal(*lastThreeBytes, []byte{'m', 'u', 'l', '('}) {
		return true
	}
	if len(*lastThreeBytes) >= 4 {
		*lastThreeBytes = (*lastThreeBytes)[1:]
	}
	return false
}

func parseMultiplyNumbers(maxLen uint8, separator byte, endChar byte, byteChunk []byte) ([]int, error) {
	numericLength := uint8(0)
	var number uint16
	var numbers []int
	for _, memoryValue := range byteChunk {
		if byteIsNumeric(memoryValue) == false {
			if memoryValue == separator && number > 0 {
				numbers = append(numbers, int(number))
				number = 0
				numericLength = 0
				continue
			}
			if memoryValue == endChar && len(numbers) == 1 {
				numbers = append(numbers, int(number))
				number = 0
				numericLength = 0
				fmt.Printf("%s == %+v\n", string(byteChunk), numbers)
				return numbers, nil
			}
			switch memoryValue {
			case 'm', 'u', 'l', '(', ')':
				continue
			default:
				return nil, fmt.Errorf("invalid char found '%s' in '%s'", string(memoryValue), string(byteChunk))
			}
		}
		numericLength++
		if numericLength > maxLen {
			return nil, fmt.Errorf("invalid multiply instruction: number too long")
		}

		var a uint16
		if memoryValue < '0' || memoryValue > '9' {
			panic("invalid character in input")
		}
		a = a*10 + uint16(memoryValue) - '0'

		if numericLength > 1 {
			number = number*10 + a
		} else {
			number = a
		}
	}

	return nil, fmt.Errorf("invalid multiply instruction: unknown state '%s'", string(byteChunk))
}

func byteIsNumeric(value byte) bool {
	switch value {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return true
	default:
		return false
	}
}

func validMemoryValue(value byte) bool {
	switch value {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'm', 'u', 'l', '(', ',', ')':
		return true
	default:
		return false
	}
}
