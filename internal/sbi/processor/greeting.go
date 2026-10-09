package processor

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p *Processor) Greet(c *gin.Context, name string) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello " + name + "!",
	})
}
