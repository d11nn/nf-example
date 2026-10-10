package processor

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p *Processor) FindMarvelCharacterName(c *gin.Context, targetName string) {
	if lastName, ok := p.Context().MarvelData[targetName]; ok {
		c.String(http.StatusOK, fmt.Sprintf("[%s] %s\n", targetName, lastName))
		return
	}
	c.String(http.StatusNotFound, fmt.Sprintf("[%s] not found in Marvel Universe!\n", targetName))
}

func (p *Processor) InsertMarvelCharacterName(c *gin.Context, charName string, realName string) {
	p.Context().MarvelData[charName] = realName

	c.String(http.StatusCreated, 
		fmt.Sprintf("[%s] %s added successfully!\n", charName, realName),
	)
}
