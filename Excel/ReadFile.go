package excel

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"

	model "github.com/AlejandroWaiz/PokemonsCardCreater/Model"
	excelize "github.com/xuri/excelize/v2"
)

var arrayOfCreatedPokemons []model.CreatedPokemon
var createdPokemon model.CreatedPokemon

func (e *Excel) ReadFile() (createdPokemons []model.CreatedPokemon, allErrors []error) {

	excelName := os.Getenv("excel_name")

	file, err := excelize.OpenFile(excelName)

	if err != nil {

		allErrors = append(allErrors, fmt.Errorf("[ReadFile | OpenFile] => %v", err))

		return nil, allErrors
	}

	sheet := os.Getenv("sheet_name")

	allRows, err := file.GetRows(sheet)

	if err != nil {

		allErrors = append(allErrors, fmt.Errorf("[Read | GetRows] => %v", err))

		return nil, allErrors

	}

	for i, onlyColumns := range allRows {

		if i == 0 {
			continue
		}

		for columnPosition, columnValue := range onlyColumns {

			if err := createPokemon(columnPosition, columnValue); err != nil {

				allErrors = append(allErrors, fmt.Errorf("[Excel Loop] => %v", err))

			}

		}

		arrayOfCreatedPokemons = append(arrayOfCreatedPokemons, createdPokemon)
		createdPokemon.Reset()

	}

	return arrayOfCreatedPokemons, nil

}

func createPokemon(columnPosition int, columnValue string) error {

	switch columnPosition {

	case 0:

		id, err := strconv.Atoi(columnValue)

		if err != nil {
			return fmt.Errorf("[createPokemon] error: %v", err)
		}

		createdPokemon.ID = id

	case 1:

		createdPokemon.Name = columnValue

		log.Printf("Creating this pokemon: %v", createdPokemon.Name)

	case 2:

		createdPokemon.GrowthRatio = columnValue

	case 3:

		hpValue, err := strconv.Atoi(columnValue)

		if err != nil {
			return fmt.Errorf("[createPokemon | HP] error: %v", err)
		}

		createdPokemon.HP = createHPPhases(hpValue, createdPokemon.GrowthRatio)

	case 4:

		attack, err := strconv.Atoi(columnValue)

		if err != nil {
			return fmt.Errorf("[createPokemon | Attack] error: %v", err)
		}

		createdPokemon.Attack = createNonHpPhases(attack, createdPokemon.GrowthRatio)

	case 5:

		defense, err := strconv.Atoi(columnValue)

		if err != nil {
			return fmt.Errorf("[createPokemon | Defense] error: %v", err)
		}

		createdPokemon.Defense = createNonHpPhases(defense, createdPokemon.GrowthRatio)

	case 6:

		spAttack, err := strconv.Atoi(columnValue)

		if err != nil {
			return fmt.Errorf("[createPokemon | SpAttack] error: %v", err)
		}

		createdPokemon.SpAttack = createNonHpPhases(spAttack, createdPokemon.GrowthRatio)

	case 7:

		spDefense, err := strconv.Atoi(columnValue)

		if err != nil {
			return fmt.Errorf("[createPokemon | SpDefense] error: %v", err)
		}

		createdPokemon.SpDefense = createNonHpPhases(spDefense, createdPokemon.GrowthRatio)

	case 8:

		speed, err := strconv.Atoi(columnValue)

		if err != nil {
			return fmt.Errorf("[createPokemon | Speed] error: %v", err)
		}

		createdPokemon.Speed = createNonHpPhases(speed, createdPokemon.GrowthRatio)

	}

	return nil

}

func createHPPhases(statValue int, GrowthRatio string) []int {

	statPhases := make([]int, 10, 10)

	for i := 0; i < 10; i++ {

		hp := float64((((2 * statValue) + 31) * ((i * 10) + 10) / 100) + (i*10 + 10) + 10)

		finalHp := int(math.Round(hp / 5))

		switch GrowthRatio {
		case "Rápido":
			if i == 1 {
				finalHp += 2
			} else if i == 0 || i == 2 {
				finalHp += 1
			}
		case "Medio":
			if i == 3 {
				finalHp += 2
			} else if i == 2 || i == 4 {
				finalHp += 1
			}
		case "Parabólico":
			if i == 5 {
				finalHp += 2
			} else if i == 4 || i == 6 {
				finalHp += 1
			}
		case "Lento":
			if i == 7 {
				finalHp += 2
			} else if i == 6 || i == 8 {
				finalHp += 1
			}
		}

		statPhases[i] = finalHp

	}

	return statPhases

}

func createNonHpPhases(statValue int, GrowthRatio string) []int {

	statPhases := make([]int, 10, 10)

	for i := 0; i < 10; i++ {

		nonHP := float64((((2 * statValue) + 31) * ((i * 10) + 10) / 100) + 5)

		finalNonHp := int(math.Round(nonHP / 10))

		switch GrowthRatio {
		case "Rápido":
			if i == 1 {
				finalNonHp += 2
			} else if i == 0 || i == 2 {
				finalNonHp += 1
			}
		case "Medio":
			if i == 3 {
				finalNonHp += 2
			} else if i == 2 || i == 4 {
				finalNonHp += 1
			}
		case "Parabólico":
			if i == 5 {
				finalNonHp += 2
			} else if i == 4 || i == 6 {
				finalNonHp += 1
			}
		case "Lento":
			if i == 7 {
				finalNonHp += 2
			} else if i == 6 || i == 8 {
				finalNonHp += 1
			}
		}

		statPhases[i] = finalNonHp

	}

	return statPhases

}
