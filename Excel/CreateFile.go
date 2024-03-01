package excel

import (
	"fmt"
	"log"
	"os"
	"strings"

	model "github.com/AlejandroWaiz/PokemonsCardCreater/Model"
	"github.com/xuri/excelize/v2"
)

func (e *Excel) CreateFile(createdPokemons []model.CreatedPokemon) (f *excelize.File, errors []error) {

	f = excelize.NewFile()

	index, err := f.NewSheet(os.Getenv("create_sheet"))

	if err != nil {

		errors = append(errors, fmt.Errorf("[CreateFile | NewSheet] error: %v", err))

	}

	errors = setDefaultRow(f, errors)

	for i, pokemon := range createdPokemons {

		for j := range model.ExcelRow {

			err := setCellValue(f, i, j, pokemon)

			if err != nil {
				errors = append(errors, fmt.Errorf("[CreateFile] => %v", err))
			}

		}

	}

	f.SetActiveSheet(index)

	return f, errors

}

func setCellValue(f *excelize.File, pokemonIndex, excelIndex int, p model.CreatedPokemon) error {

	switch excelIndex {

	case 0:

		cell := fmt.Sprintf("A%v", pokemonIndex+2)

		err := f.SetCellValue(os.Getenv("create_sheet"), cell, p.ID)

		if err != nil {
			return fmt.Errorf("[setCellValue] error: %v", err)
		}

	case 1:

		cell := fmt.Sprintf("B%v", pokemonIndex+2)

		err := f.SetCellValue(os.Getenv("create_sheet"), cell, p.Name)

		if err != nil {
			return fmt.Errorf("[setCellValue] error: %v", err)
		}

	case 2:

		cell := fmt.Sprintf("C%v", pokemonIndex+2)

		err := f.SetCellValue(os.Getenv("create_sheet"), cell, p.GrowthRatio)

		if err != nil {
			return fmt.Errorf("[setCellValue] error: %v", err)
		}

	case 3:

		cell := fmt.Sprintf("D%v", pokemonIndex+2)

		arr := arrayToString(p.HP, "|")

		err := f.SetCellValue(os.Getenv("create_sheet"), cell, arr)

		if err != nil {
			return fmt.Errorf("[setCellValue] error: %v", err)
		}

	case 4:
		cell := fmt.Sprintf("E%v", pokemonIndex+2)

		arr := arrayToString(p.Attack, "|")

		err := f.SetCellValue(os.Getenv("create_sheet"), cell, arr)

		if err != nil {
			return fmt.Errorf("[setCellValue] error: %v", err)
		}

	case 5:

		cell := fmt.Sprintf("F%v", pokemonIndex+2)

		arr := arrayToString(p.Defense, "|")

		err := f.SetCellValue(os.Getenv("create_sheet"), cell, arr)

		if err != nil {
			return fmt.Errorf("[setCellValue] error: %v", err)
		}

	case 6:
		cell := fmt.Sprintf("G%v", pokemonIndex+2)

		arr := arrayToString(p.SpAttack, "|")

		err := f.SetCellValue(os.Getenv("create_sheet"), cell, arr)

		if err != nil {
			return fmt.Errorf("[setCellValue] error: %v", err)
		}

	case 7:

		cell := fmt.Sprintf("H%v", pokemonIndex+2)

		arr := arrayToString(p.SpDefense, "|")

		err := f.SetCellValue(os.Getenv("create_sheet"), cell, arr)

		if err != nil {
			return fmt.Errorf("[setCellValue] error: %v", err)
		}

	case 8:

		cell := fmt.Sprintf("I%v", pokemonIndex+2)

		arr := arrayToString(p.Speed, "|")

		err := f.SetCellValue(os.Getenv("create_sheet"), cell, arr)

		if err != nil {
			return fmt.Errorf("[setCellValue] error: %v", err)
		}
	}

	return nil
}

func arrayToString(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
	//return strings.Trim(strings.Join(strings.Split(fmt.Sprint(a), " "), delim), "[]")
	//return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(a)), delim), "[]")
}

func setDefaultRow(f *excelize.File, errors []error) []error {

	for i, cellName := range model.ExcelRow {

		coord := fmt.Sprintf("%v%v", model.ExcelCells[i], 1)

		err := f.SetCellValue(os.Getenv("create_sheet"), coord, cellName)

		if err != nil {

			log.Println(err)

			errors = append(errors, fmt.Errorf("[setDefaultRow | setCellValue] error: %v", err))

		}

	}

	return errors
}
