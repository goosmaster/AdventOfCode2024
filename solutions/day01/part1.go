package day01

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
	"sort"
)

func Part1() (string, error) {
	file, err := os.OpenFile("solutions/day01/input.txt", os.O_RDONLY, 0600)
	if err != nil {
		return "", err
	}

	answer, puzzleErr := solvePart1(bufio.NewReader(file))
	return fmt.Sprintf("%d", answer), puzzleErr
}

func solvePart1(file *bufio.Reader) (uint32, error) {
	listScanner := bufio.NewScanner(file)
	lineIndex := uint16(0)
	leftList := make(map[uint16]uint32, 1000)
	rightList := make(map[uint16]uint32, 1000)

	for listScanner.Scan() {
		lists := listScanner.Bytes()
		listParts := bytes.Split(lists, []byte{' ', ' ', ' '})

		appendPart(lineIndex, &listParts[0], &leftList)
		appendPart(lineIndex, &listParts[1], &rightList)

		if lineIndex == math.MaxUint16 {
			return 0, fmt.Errorf("reached end of uint16")
		}
		lineIndex++
	}

	leftList = sortList(leftList)
	rightList = sortList(rightList)

	total := uint32(0)
	for index, leftValue := range leftList {
		total += diff(leftValue, rightList[index])
	}

	return total, nil
}

func sortList(list map[uint16]uint32) map[uint16]uint32 {
	result := make(map[uint16]uint32, len(list))
	keys := make([]uint16, 0, len(list))
	for key := range list {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool {
		return list[keys[i]] > list[keys[j]]
	})
	for index, key := range keys {
		result[uint16(index)] = list[key]
	}
	return result
}

func diff(a, b uint32) uint32 {
	if a > b {
		return a - b
	}
	return b - a
}

func appendPart(lineIndex uint16, listPart *[]byte, targetList *map[uint16]uint32) {
	var number uint32
	for _, b := range *listPart {
		number = number*10 + uint32(b-'0')
	}

	(*targetList)[lineIndex] = number
}
