package day02

import (
	"bufio"
	"bytes"
	"testing"
)

func TestPart1(t *testing.T) {
	const input = `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`
	const expectedResult = 2
	result, err := solvePart1(bufio.NewReader(bytes.NewBuffer([]byte(input))))
	if err != nil {
		t.Errorf("got err: %v", err)
	}
	if result != expectedResult {
		t.Errorf("got %d, expected %d", result, expectedResult)
	}
}
