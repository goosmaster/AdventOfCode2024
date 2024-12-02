package day02

import (
	"bufio"
	"bytes"
	"testing"
)

func TestPart2(t *testing.T) {
	// "better" test cases from AoC Reddit
	//[57] 56 57 59 60 63 64 65
	//91 92 [95] 93 94
	//16 [13] 15 13 12 11 9 6
	//40 41 43 44 [47] 46 47 49
	//	const input = `57 56 57 59 60 63 64 65
	//91 92 95 93 94
	//16 13 15 13 12 11 9 6
	//40 41 43 44 47 46 47 49
	//12 10 13 16 19 21 22`

	const input = `48 46 47 49 51 54 56
1 1 2 3 4 5
1 2 3 4 5 5
5 1 2 3 4 5
1 4 3 2 1
1 6 7 8 9
1 2 3 4 3
9 8 7 6 7
7 10 8 10 11
29 28 27 25 26 25 22 20`

	const expectedResult = 10
	result, err := solvePart2(bufio.NewReader(bytes.NewBuffer([]byte(input))))
	if err != nil {
		t.Errorf("got err: %v", err)
	}
	if result != expectedResult {
		t.Errorf("got %d, expected %d", result, expectedResult)
	}
}
