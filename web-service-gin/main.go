package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	router := gin.Default()
	// relaciona a rota /albums com a função getAlbums
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumById)
	router.DELETE("/albums/:id", deleteAlbumById)
	router.POST("/albums", postAlbums)
	// sobe o servidor
	router.Run("localhost:8080")
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	// retorna ao frontend a lista de album no formato JSON
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(frontend *gin.Context) {
	var vinil album
	// vinil vai guardar os dados do frontend
	if err := frontend.BindJSON(&vinil); err != nil {
		return // tem erro
	}
	// adiciona vinil no album
	albums = append(albums, vinil)
	frontend.IndentedJSON(http.StatusCreated, vinil)
}

func getAlbumById(frontend *gin.Context) {
	id := frontend.Param("id")
	for _, a := range albums {
		if a.ID == id {
			frontend.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	frontend.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func deleteAlbumById(frontend *gin.Context) {
	id := frontend.Param("id")
	// cria um vetor auxiliar
	aux := []album{}

	for _, a := range albums {
		if a.ID != id {
			aux = append(aux, a)
		}
	}
	if len(aux) == len(albums) {
		frontend.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	}
	albums = aux
	frontend.IndentedJSON(http.StatusOK, gin.H{"message": "album deleted"})

}
