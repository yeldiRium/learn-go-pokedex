package model

type Pokemon struct {
	Name           string
	BaseExperience uint
	Height         uint
	Weight         uint
	Abilities      []string
	Species        string
	BaseStats      Stats
	Types          []string
}
