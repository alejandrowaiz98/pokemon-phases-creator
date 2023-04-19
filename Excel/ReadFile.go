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

		hpValue, err := strconv.Atoi(columnValue)

		if err != nil {
			return fmt.Errorf("[createPokemon | HP] error: %v", err)
		}

		createdPokemon.HP = createHPPhases(hpValue)

	case 3:

		attack, err := strconv.Atoi(columnValue)

		if err != nil {
			return fmt.Errorf("[createPokemon | Attack] error: %v", err)
		}

		createdPokemon.Attack = createNonHpPhases(attack)

	case 4:

		defense, err := strconv.Atoi(columnValue)

		if err != nil {
			return fmt.Errorf("[createPokemon | Defense] error: %v", err)
		}

		createdPokemon.Defense = createNonHpPhases(defense)

	case 5:

		spAttack, err := strconv.Atoi(columnValue)

		if err != nil {
			return fmt.Errorf("[createPokemon | SpAttack] error: %v", err)
		}

		createdPokemon.SpAttack = createNonHpPhases(spAttack)

	case 6:

		spDefense, err := strconv.Atoi(columnValue)

		if err != nil {
			return fmt.Errorf("[createPokemon | SpDefense] error: %v", err)
		}

		createdPokemon.SpDefense = createNonHpPhases(spDefense)

	case 7:

		speed, err := strconv.Atoi(columnValue)

		if err != nil {
			return fmt.Errorf("[createPokemon | Speed] error: %v", err)
		}

		createdPokemon.Speed = createNonHpPhases(speed)

	}

	return nil

}

func createHPPhases(statValue int) []int {

	statPhases := make([]int, 10, 10)

	for i := 0; i < 10; i++ {

		var hp float64

		hp = float64((((2 * statValue) + 31) * ((i * 10) + 10) / 100) + (i*10 + 10) + 10)

		statPhases[i] = int(math.Round(hp / 5))

	}

	return statPhases

}

func createNonHpPhases(statValue int) []int {

	statPhases := make([]int, 10, 10)

	for i := 0; i < 10; i++ {

		var nonHP float64

		nonHP = float64((((2 * statValue) + 31) * ((i * 10) + 10) / 100) + 5)

		statPhases[i] = int(math.Round(nonHP / 10))

	}

	return statPhases

}
