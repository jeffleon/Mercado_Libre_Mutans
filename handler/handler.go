package handler

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/jeffleon/mutansApi/database"
	"github.com/jeffleon/mutansApi/model"
	"github.com/jeffleon/mutansApi/utils"
	_ "github.com/jeffleon/mutansApi/utils"
)

type Mutant struct {
	Dna []string `json:"dna"`
}
func PostMutant(c *fiber.Ctx) error {
	var DNAmutanType Mutant
	var mutant model.Mutants
	var arrayDNA []string
	db := database.DBConn

	DNAmutan := c.Body()
	err := json.Unmarshal([]byte(DNAmutan), &DNAmutanType)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't do a json Unmarshall", "data": err})
	}
	arrayDNA = DNAmutanType.Dna
	if len(arrayDNA) == 0 {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "json data error"})
	}
	isMutant := utils.IsMutant(arrayDNA)
	if len(isMutant) == 0 {
		return c.Status(403).JSON(fiber.Map{"status": "error", "message": "Forbiden DNA user"})
	}
	mutant.CountHumanDna = 100
	mutant.CountMutantDna = len(isMutant)
	mutant.Ratio = float32(len(isMutant)) * 0.1
	db.Create(&mutant)
	return c.Status(200).JSON(fiber.Map{"status": "success", "data": "this DNA contains mutant gene"})
}

func GetStats(c *fiber.Ctx) error {
	var mutant model.Mutants
	id := c.Params("id")
	db := database.DBConn
	db.First(&mutant, id)

	return c.Status(200).JSON(fiber.Map{"status": "success", "data": mutant})
}
