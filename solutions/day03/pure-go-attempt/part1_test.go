package day03

import (
	"bufio"
	"bytes"
	"testing"
)

func TestPart1(t *testing.T) {
	//const input = `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`

	// got 840115, expected 1384648
	//const input = `!?~[ don't():mul(850,510))[#}how()mul(339,21){[select()@what()#(*mul(850,794)where()}where()mul(732;mul(949,117)<-why()@ mul(164,964)why()where()do()>; w`
	const input = `aaa!?~[ don't():mul(850,510))[#}how()mul(339,21){[select()@what()#(*mul(850,794)where()}where()mul(732;mul(949,117)<-why()@ mul(164,964)why()where()do()>; w`
	const expectedResult = 1384648
	result, err := solvePart1(bufio.NewReader(bytes.NewBuffer([]byte(input))))
	if err != nil {
		t.Errorf("got err: %v", err)
	}
	if result != expectedResult {
		t.Errorf("got %d, expected %d", result, expectedResult)
	}
}
