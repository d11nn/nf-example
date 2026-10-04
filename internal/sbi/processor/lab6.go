package processor

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p *Processor) HelloLab6(c *gin.Context, name string) {
	if name == "" {
		c.String(http.StatusBadRequest, "No name provided")
		return
	}

	c.String(
		http.StatusOK,
		fmt.Sprintf("Hello %s, welcome to free5GC Lab6!", name),
	)
}
func (p *Processor) PostHelloLab6(c *gin.Context, name string, message string) {
	if name == "" || message == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "name and message are required",
		})
		return
	}

	result := fmt.Sprintf("Hello %s! %s", name, message)

	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}
