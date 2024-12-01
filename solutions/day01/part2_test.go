package day01

import (
	"bufio"
	"bytes"
	"testing"
)

func TestPart2(t *testing.T) {
	const input = `3   4
4   3
2   5
1   3
3   9
3   3`

	const expectedResult = uint32(31)
	result, err := solvePart2(bufio.NewReader(bytes.NewBuffer([]byte(input))))
	if err != nil {
		t.Errorf("got err: %v", err)
	}
	if result != expectedResult {
		t.Errorf("got %d, expected %d", result, expectedResult)
	}
}
