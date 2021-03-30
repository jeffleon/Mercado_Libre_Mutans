package utils

type Map = map[string]interface{}

func IsMutant(dna []string) []map[string]interface{}{
	var DNAMutant []map[string]interface{}
	var i, j int


	for i = 0; i < len(dna[0]); i++ {

		for j = 0; j < len(dna); j++{

			if (len(dna[0]) - i) >= 4 {
				if dna[i][j] == dna[i + 1][j]{
					validation(string(dna[i][j]), dna, i, j, "Down", &DNAMutant)
				}
				if j < (len(dna) - 1) && i < (len(dna[0]) - 1){
					if dna[i][j] == dna[i + 1][j + 1]{
						validation(string(dna[i][j]), dna, i, j, "Lateral", &DNAMutant)
					}
				}
			}
			if (len(dna) - j) > 1 {
				if dna[i][j] == dna[i][j + 1]{
					validation(string(dna[i][j]), dna, i, j, "Right", &DNAMutant)
				}
			}
		}
	}

	return DNAMutant
}


func validation(letter string,matrix []string, row int, column int, comand string, DNAMutant *[]map[string]interface{}) bool{
	var i, j, count int = 0, 0 , 0
	var axisX ,axisY int = 0, 0

	if comand == "Lateral" || comand == "Right"{
		axisX = 1
	}

	if comand == "Lateral" || comand == "Down"{
		axisY = 1
	}
	if comand != "Down"{

	}
	for i = row; i < len(matrix[0]); {
		for j = column; j < len(matrix); {

			if (j < len(matrix) - 1) && (i <= len(matrix[0]) - 1) && matrix[i][j] == matrix[i + axisY][j + axisX] {
				if comand == "Down"{
					i++
				} else if comand == "Lateral"{
					i++
					j++
				} else {
					j++
				}
				count++
			} else {
				return false
			}

			if count == 3 {
				*DNAMutant = append(*DNAMutant, Map{"mutant": true,"letter": letter, "axisx":row, "axisy": column, "side": comand})
				return true
			}
		}
	}
	return true
}




