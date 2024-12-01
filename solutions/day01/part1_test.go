package day01

import (
	"bufio"
	"bytes"
	"testing"
)

func TestPart1(t *testing.T) {
	const input = `37033   48086
80098   34930
88073   69183
54342   63061
98409   87908
81400   96222
42062   53621
55208   48086
10847   20622
53237   11766
12609   19507
31524   33054
83455   96879
53344   76641
94982   66380
69183   70224
35580   12846`

	const expectedResult = uint32(92868)
	result, err := solvePart1(bufio.NewReader(bytes.NewBuffer([]byte(input))))
	if err != nil {
		t.Errorf("got err: %v", err)
	}
	if result != expectedResult {
		t.Errorf("got %d, expected %d", result, expectedResult)
	}
}
