package sbi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Pokemon struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

var pokemonData = map[string]string{
	"Pikachu":    "Electric",
	"Charmander": "Fire",
	"Squirtle":   "Water",
	"Bulbasaur":  "Grass",
}

func (s *Server) getPokemonRoute() []Route {
	return []Route{
		{
			Name:    "Pokemon Home",
			Method:  http.MethodGet,
			Pattern: "/",
			APIFunc: s.HTTPPokemonHome,
		},
		{
			Name:    "Get Pokemon",
			Method:  http.MethodGet,
			Pattern: "/:Name",
			APIFunc: s.HTTPGetPokemon,
		},
		{
			Name:    "Create Pokemon",
			Method:  http.MethodPost,
			Pattern: "/",
			APIFunc: s.HTTPCreatePokemon,
		},
	}
}

func (s *Server) HTTPPokemonHome(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello Pokemon!")
}

func (s *Server) HTTPGetPokemon(c *gin.Context) {
	name := c.Param("Name")

	pokemonType, exists := pokemonData[name]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "pokemon not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"name": name,
		"type": pokemonType,
	})
}

func (s *Server) HTTPCreatePokemon(c *gin.Context) {
	var pokemon Pokemon

	if err := c.ShouldBindJSON(&pokemon); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	pokemonData[pokemon.Name] = pokemon.Type

	c.JSON(http.StatusCreated, gin.H{
		"name": pokemon.Name,
		"type": pokemon.Type,
	})
}
