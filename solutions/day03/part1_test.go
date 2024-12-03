package day03

import (
	"bufio"
	"bytes"
	"testing"
)

func TestPart1(t *testing.T) {
	const input = `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`

	const expectedResult = 161
	result, err := solvePart1(bufio.NewReader(bytes.NewBuffer([]byte(input))))
	if err != nil {
		t.Errorf("got err: %v", err)
	}
	if result != expectedResult {
		t.Errorf("got %d, expected %d", result, expectedResult)
	}
}
