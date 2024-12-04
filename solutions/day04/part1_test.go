package day04

import (
	"bufio"
	"bytes"
	"testing"
)

func TestPart1(t *testing.T) {
	testCases := []struct {
		name           string
		input          string
		expectedResult uint32
	}{
		{"horizontally", "SAMXMAS", 2},
		{"horizontally reversed", "SAMXXMAS", 2},
		{"vertically", "X.S\nM.A\nA.M\nS.X", 2},
		{"vertically reversed", "S..\nA..\nM..\nX..", 1},
		{"diagonally", "X...\n.M..\n..A.\n...S", 1},
		{"diagonally reversed", "S...\n.A..\n..M.\n...X", 1},
		{"Top left", "XMAS\nMM..\nA.A.\nS..S", 3},
		{"Test input", "MMMSXXMASM\nMSAMXMSMSA\nAMXSXMAAMM\nMSAMASMSMX\nXMASAMXAMM\nXXAMMXXAMA\nSMSMSASXSS\nSAXAMASAAA\nMAMMMXMMMM\nMXMXAXMASX", 18},
		{"example 1st", "....X.....\n.....M....\n......A...\n.......S..\n..........", 1},
		{"example 2nd", ".....XMAS.\n..........", 1},
		{"example 3rd", "..........\n.SAMX.....\n..........", 1},
		{"example 4th", "S.S.S.S.SS\n.A.A.A.A.A\n..M.M.M.MM\n.X.X.XMASX", 8},
		{"example 5th", "XMASAMX.MM", 2},
		{"example 6th", ".S.....S..\n..A...A...\n...M.M....\n....X.....\n...M.M....\n..A...A...\n.S.....S..", 4},
	}

	t.Parallel()
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result, err := solvePart1(bufio.NewReader(bytes.NewBuffer([]byte(testCase.input))))
			if err != nil {
				t.Errorf("got err: %v", err)
			}
			if result != testCase.expectedResult {
				t.Errorf("got %d, expected %d", result, testCase.expectedResult)
			}
		})
	}
}
