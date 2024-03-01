package model

type PokemonToCreate struct {
	Name, GrowthRatio                                   string
	ID, HP, Attack, Defense, SpAttack, SpDefense, Speed int
}

type CreatedPokemon struct {
	//All arrays represent the 10 phases of each pokemon
	Name, GrowthRatio                               string
	ID                                              int
	HP, Attack, Defense, SpAttack, SpDefense, Speed []int
}

var emptyPokemon = &CreatedPokemon{}

func (p *CreatedPokemon) Reset() {

	*p = *emptyPokemon

}
