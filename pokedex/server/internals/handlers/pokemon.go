package handlers

import (
	"pokedex/internals/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPokeDex(c *gin.Context) {

	value, err := services.FetchPokedex()
	if err != nil {
		c.JSON(500, gin.H{
			"message": "failed to fetch pokedex",
		})
		return
	}

	c.JSON(200, gin.H{
		"data": value,
	})

}

func GetPokemon(c *gin.Context) {

	offset, err := strconv.Atoi(c.Query("offset"))
	if err != nil {
		offset = 0
	}

	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		limit = 10
	}

	opts := &services.FetchPokemonOpts{
		Offset: offset,
		Limit:  limit,
	}

	resource, err := services.FetchPokemon(opts)

	if err != nil {
		c.JSON(500, gin.H{
			"message": "failed to fetch pokemon",
		})
		return
	}

	c.JSON(200, gin.H{
		"data": resource,
	})

}

func GetPokemonByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(400, gin.H{
			"message": "invalid id",
		})
		return
	}

	resource, err := services.FetchPokemonByID(id)

	if err != nil {
		c.JSON(500, gin.H{
			"message": "failed to fetch pokemon",
		})
		return
	}

	c.JSON(200, gin.H{
		"data": resource,
	})

}

func SearchPokemon(c *gin.Context) {
	identifier := c.Query("identifier")

	resource, err := services.SearchPokemon(identifier)

	if err != nil {
		c.JSON(500, gin.H{
			"message": "failed to search pokemon",
		})
		return
	}

	c.JSON(200, gin.H{
		"data": resource,
	})

}
