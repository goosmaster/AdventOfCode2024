package day03

import (
	"bufio"
	"bytes"
	"testing"
)

func TestPart2(t *testing.T) {
	const input = `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`

	const expectedResult = 48
	result, err := solvePart2(bufio.NewReader(bytes.NewBuffer([]byte(input))))
	if err != nil {
		t.Errorf("got err: %v", err)
	}
	if result != expectedResult {
		t.Errorf("got %d, expected %d", result, expectedResult)
	}
}
