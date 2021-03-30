package tests

import (
	"github.com/jeffleon/mutansApi/utils"
	"testing"
)

var mutants = []string{ "ATGCGC",
						"CAGTGC",
						"TTATAA",
						"AGAAGC",
						"CCCATA",
						"TTTTCG"}

var mutants1 = []string{ "ATGCGC",
						 "CAGTGC",
						 "TTATAA",
						 "AGAGGC",
						 "CCCATA",
						 "TTTACG"}


var mutants2 = []string{ "AAAATC",
						 "CAGTGC",
						 "TTATAC",
						 "AGAACA",
						 "CCCCTA",
						 "GGGGAT"}

var mutants3 = []string{ "ATGCGC",
						 "CAGTGT",
						 "TTATAG",
						 "AGAAGC",
						 "CCCATA",
						 "AAACCG"}

type isMutant struct {
	arg1 []string
	expected int
}



var MutantTest = []isMutant{
	isMutant{mutants, 2},
	isMutant{mutants1, 0},
	isMutant{mutants2, 4},
	isMutant{mutants3, 1},
}

func TestAdd(t *testing.T){

	for _, test := range MutantTest {
		if output := utils.IsMutant(test.arg1); len(output) != test.expected {

			t.Errorf("Output %q not equal to expected %q", len(output), test.expected)
		}
	}
}


