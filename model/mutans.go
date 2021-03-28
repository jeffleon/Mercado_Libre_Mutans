package model

import (
	"github.com/jinzhu/gorm"
)

type Mutants struct {
	gorm.Model
	CountMutantDna int `json:"count_mutant_dna"`
	CountHumanDna int `json:"count_human_dna"`
	Ratio float32 `json:"ratio"`
}
