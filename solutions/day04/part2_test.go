package day04

import (
	"bufio"
	"bytes"
	"testing"
)

func TestPart2(t *testing.T) {
	testCases := []struct {
		name           string
		input          string
		expectedResult uint32
	}{
		{"example", ".M.S......\n..A..MSMS.\n.M.S.MAA..\n..A.ASMSM.\n.M.S.M....\n..........\nS.S.S.S.S.\n.A.A.A.A..\nM.M.M.M.M.\n..........", 9},
		{"mini example", "M.S\n.A.\nM.S", 1},
	}

	t.Parallel()
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result, err := solvePart2(bufio.NewReader(bytes.NewBuffer([]byte(testCase.input))))
			if err != nil {
				t.Errorf("got err: %v", err)
			}
			if result != testCase.expectedResult {
				t.Errorf("got %d, expected %d", result, testCase.expectedResult)
			}
		})
	}
}
