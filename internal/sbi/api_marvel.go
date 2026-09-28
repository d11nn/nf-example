package sbi

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func (s *Server) HTTPSearchMarvelCharacter(c *gin.Context) {
	targetName := c.Param("Name")

	if targetName == "" {
		c.String(http.StatusBadRequest, "No name provided!")
		return
	}

	s.Processor().FindMarvelCharacterName(c, targetName)
}

func (s *Server) HTTPAddMarvelCharacter(c *gin.Context) {
	type RequestBody struct {
		Name     string `json:"name"`
		RealName string `json:"realName"`
	}

	var body RequestBody

	if err := c.ShouldBindJSON(&body); err != nil {
		c.String(http.StatusBadRequest, "Invalid request body")
		return
	}

	if body.Name == "" || body.RealName == "" {
		c.String(http.StatusBadRequest, "Both character name and real name are required")
		return
	}

	s.Processor().InsertMarvelCharacterName(c, body.Name, body.RealName)
}

func (s *Server) getMarvelRoute() []Route {
    return []Route{
        {
            Name:    "Hello Marvel Universe!",
            Method:  http.MethodGet,
            Pattern: "/",
            APIFunc: func(c *gin.Context) {
                c.JSON(http.StatusOK, "Hello Marvel Universe!")
            },
        },
		{
			Name:    "Get Marvel Characters' Name.",
			Method:  http.MethodGet,
			Pattern: "/character/:Name",
			APIFunc: s.HTTPSearchMarvelCharacter,
		},
        {
            Name:    "Create Marvel Character.",
            Method:  http.MethodPost,
            Pattern: "/character",
            APIFunc: s.HTTPAddMarvelCharacter,
        },
    }
}