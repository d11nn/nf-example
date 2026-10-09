package sbi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) getGreetingRoute() []Route {
	return []Route{
		{
			Name:    "Get greeting",
			Method:  http.MethodGet,
			Pattern: "/",
			APIFunc: s.HTTPGetGreeting,
		},
		{
			Name:    "Post greeting",
			Method:  http.MethodPost,
			Pattern: "/",
			APIFunc: s.HTTPPostGreeting,
		},
	}
}

func (s *Server) HTTPGetGreeting(c *gin.Context) {
	name := c.DefaultQuery("name", "free5GC")
	s.Processor().Greet(c, name)
}

func (s *Server) HTTPPostGreeting(c *gin.Context) {
	var request struct {
		Name string `json:"name" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "name is required; send a JSON body",
		})
		return
	}

	s.Processor().Greet(c, request.Name)
}
