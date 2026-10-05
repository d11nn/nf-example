package processor

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p *Processor) Findmessagetext(c *gin.Context, targetname string) {
	if food, ok := p.Context().MessageData[targetname]; ok {
		c.String(http.StatusOK, fmt.Sprintf("food: %s %s", targetname, food))
		return
	}
	c.String(http.StatusNotFound, fmt.Sprintf("[%s] not found in MessageData", targetname))
}



func (p *Processor) Insert(c *gin.Context,name string,food string) {
	p.Context().MessageData[name]=food
	
	c.JSON(http.StatusCreated, gin.H{
    "name": name,
    "food": food,
	})
}
