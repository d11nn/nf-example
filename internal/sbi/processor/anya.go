package processor

import (
	"net/http"

	"github.com/Alonza0314/nf-example/internal/logger"
	"github.com/gin-gonic/gin"
)

func (p *Processor) FindAnyaCharacterName(c *gin.Context, anyaName string) {
	logger.SBILog.Infof("FindAnyaCharacterName: %s", anyaName)
	c.JSON(http.StatusOK, gin.H{"message": "Anya " + anyaName + " is here!"})
}

func (p *Processor) CreateAnyaCharacter(c *gin.Context) {
	logger.SBILog.Infof("CreateAnyaCharacter")
	c.JSON(http.StatusCreated, gin.H{"message": "Anya created!"})
}
