package model

type PokemonAbilities []string
type PokemonTypes []string

type Pokemon struct {
	Name           string
	BaseExperience uint
	Height         uint
	Weight         uint
	Abilities      PokemonAbilities
	Species        string
	BaseStats      Stats
	Types          PokemonTypes
}
