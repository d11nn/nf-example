package sbi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type chiikawaVoice struct {
	Name  string `json:"name" binding:"required"`
	Voice string `json:"voice" binding:"required"`
}

var chiikawaVoices = map[string]string{
	"Chiikawa":  "Soft whimpers, happy squeaks, and tiny crying sounds.",
	"Hachiware": "Oi, oi, oi!",
	"Momonga":   "Iya yada yada!",
	"Usagi":     "Una una una! Yaha!",
}

// getChiikawaRoute returns the route list for the Chiikawa voice service.
func (s *Server) getChiikawaRoute() []Route {
	return []Route{
		{
			Name:    "Hello Chiikawa!",
			Method:  http.MethodGet,
			Pattern: "/",
			APIFunc: func(c *gin.Context) {
				c.JSON(http.StatusOK, "Hello Chiikawa!")
			},
		},
		{
			Name:    "Chiikawa Character Voice",
			Method:  http.MethodGet,
			Pattern: "/voice/:Name",
			APIFunc: s.HTTPGetChiikawaVoice,
		},
		{
			Name:    "Create Chiikawa Character Voice",
			Method:  http.MethodPost,
			Pattern: "/voice",
			APIFunc: s.HTTPCreateChiikawaVoice,
		},
	}
}

// HTTPGetChiikawaVoice returns the saved voice for a Chiikawa character.
func (s *Server) HTTPGetChiikawaVoice(c *gin.Context) {
	targetName := c.Param("Name")
	if targetName == "" {
		c.String(http.StatusBadRequest, "No name provided")
		return
	}

	voice, ok := chiikawaVoices[targetName]
	if !ok {
		c.String(http.StatusNotFound, "%s voice not found", targetName)
		return
	}

	c.JSON(http.StatusOK, chiikawaVoice{
		Name:  targetName,
		Voice: voice,
	})
}

// HTTPCreateChiikawaVoice saves a Chiikawa character voice from JSON input.
func (s *Server) HTTPCreateChiikawaVoice(c *gin.Context) {
	var request chiikawaVoice
	if err := c.ShouldBindJSON(&request); err != nil {
		c.String(http.StatusBadRequest, "Invalid request body")
		return
	}

	chiikawaVoices[request.Name] = request.Voice
	c.JSON(http.StatusCreated, gin.H{
		"message": "Chiikawa voice saved",
		"name":    request.Name,
		"voice":   request.Voice,
	})
}
