package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"slices"

	"github.com/yeldiRium/learning-go-pokedex/model"
	"github.com/yeldiRium/learning-go-pokedex/pokecache"
)

var ErrGetPokemon = errors.New("GetPokemon")
var ErrPokemonDoesntExist = errors.New("pokemon does not exist")

const BasePokemonUrl = "https://pokeapi.co/api/v2/pokemon/"

func GetPokemon(httpClient HttpClient, cache pokecache.Cache, pokemonName string) (model.Pokemon, error) {
	url := fmt.Sprintf("%s%s/", BasePokemonUrl, pokemonName)

	responseBody, ok := cache.GetEntry(url)
	if !ok {
		request, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return model.Pokemon{}, fmt.Errorf("%w: %w: %w", ErrGetPokemon, ErrRequestInvalid, err)
		}
		response, err := httpClient.Do(request)
		if err != nil {
			return model.Pokemon{}, fmt.Errorf("%w: %w: %w", ErrGetPokemon, ErrRequestFailed, err)
		}
		if response.StatusCode == http.StatusNotFound {
			return model.Pokemon{}, fmt.Errorf("%w: %w: %s", ErrGetPokemon, ErrPokemonDoesntExist, pokemonName)
		}

		defer response.Body.Close()
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return model.Pokemon{}, fmt.Errorf("%w: %w: %w", ErrGetPokemon, ErrRequestFailed, err)
		}

		responseBody = body
	}

	var apiResponse pokemonApiResponse
	err := json.Unmarshal(responseBody, &apiResponse)
	if err != nil {
		return model.Pokemon{}, fmt.Errorf("%w: %w: %w", ErrGetPokemon, ErrRequestFailed, err)
	}

	return ParsePokemon(apiResponse)
}

type pokemonTypeApiResponse struct {
	Slot uint `json:"slot"`
	Type struct {
		Name string `json:"name"`
	} `json:"type"`
}

func ParsePokemonTypes(rawTypes []pokemonTypeApiResponse) model.PokemonTypes {
	types := make([]string, len(rawTypes))
	slices.SortFunc(rawTypes, func(a, b pokemonTypeApiResponse) int {
		return int(a.Slot) - int(b.Slot)
	})
	for i, pokemonType := range rawTypes {
		types[i] = pokemonType.Type.Name
	}
	return types
}

type pokemonAbilityApiResponse struct {
	Slot    uint `json:"slot"`
	Ability struct {
		Name string `json:"name"`
	} `json:"ability"`
}

func ParsePokemonAbilities(rawAbilities []pokemonAbilityApiResponse) model.PokemonAbilities {
	abilities := make([]string, len(rawAbilities))
	slices.SortFunc(rawAbilities, func(a, b pokemonAbilityApiResponse) int {
		return int(a.Slot) - int(b.Slot)
	})
	for i, ability := range rawAbilities {
		abilities[i] = ability.Ability.Name
	}
	return abilities
}

type pokemonStatApiResponse struct {
	BaseStat uint `json:"base_stat"`
	Stat     struct {
		Name string `json:"name"`
	} `json:"stat"`
}

func ParsePokemonStats(rawStats []pokemonStatApiResponse) (model.Stats, error) {
	stats := model.Stats{}
	for _, stat := range rawStats {
		switch stat.Stat.Name {
		case "hp":
			stats.Hp = stat.BaseStat
		case "attack":
			stats.Attack = stat.BaseStat
		case "defense":
			stats.Defense = stat.BaseStat
		case "special-attack":
			stats.SpecialAttack = stat.BaseStat
		case "special-defense":
			stats.SpecialDefense = stat.BaseStat
		case "speed":
			stats.Speed = stat.BaseStat
		default:
			return model.Stats{}, fmt.Errorf("%w: unknown stat: '%s'", ErrGetPokemon, stat.Stat.Name)
		}
	}
	return stats, nil
}

type pokemonApiResponse struct {
	Name           string                      `json:"name"`
	BaseExperience uint                        `json:"base_experience"`
	Height         uint                        `json:"height"`
	Weight         uint                        `json:"weight"`
	Abilities      []pokemonAbilityApiResponse `json:"abilities"`
	Species        struct {
		Name string `json:"name"`
	} `json:"species"`
	Stats []pokemonStatApiResponse `json:"stats"`
	Types []pokemonTypeApiResponse `json:"types"`
}

func ParsePokemon(rawPokemon pokemonApiResponse) (model.Pokemon, error) {
	types := ParsePokemonTypes(rawPokemon.Types)
	abilities := ParsePokemonAbilities(rawPokemon.Abilities)
	stats, err := ParsePokemonStats(rawPokemon.Stats)
	if err != nil {
		return model.Pokemon{}, err
	}

	pokemon := model.Pokemon{
		Name:           rawPokemon.Name,
		BaseExperience: rawPokemon.BaseExperience,
		Height:         rawPokemon.Height,
		Weight:         rawPokemon.Weight,
		Species:        rawPokemon.Species.Name,
		Types:          types,
		Abilities:      abilities,
		BaseStats:      stats,
	}

	return pokemon, nil
}
