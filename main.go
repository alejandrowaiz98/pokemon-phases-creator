package main

import (
	"log"
	"os"

	Excel "github.com/AlejandroWaiz/PokemonsCardCreater/Excel"

	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load(".env")

	log.Println(os.Getenv("excel_name"))

	excel := Excel.New()

	pokemons, errors := excel.ReadFile()

	if len(errors) > 0 {
		for _, err := range errors {
			log.Println(err)
		}
	}

	file, errors := excel.CreateFile(pokemons)

	if len(errors) > 0 {
		for _, err := range errors {
			log.Println(err)
		}
	}

	err := file.SaveAs("CreatedPokemons.xlsx")

	if err != nil {
		log.Println(err)
	}

}
