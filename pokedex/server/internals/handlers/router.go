package handlers

import "github.com/gin-gonic/gin"

func InitializeRouter(r *gin.Engine) {

	r.GET("/ping", GetPing)
	r.GET("/health", HealthCheck)
	r.GET("/pokedex", GetPokeDex)
	r.GET("/pokemon", GetPokemon)
	r.GET("/pokemon/one/:id", GetPokemonByID)
	r.GET("/pokemon/search", SearchPokemon)

}
