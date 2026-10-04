package sbi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Lab6HelloRequest struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}

func (s *Server) getLab6Route() []Route {
	return []Route{
		{
			Name:    "Lab6 Hello GET",
			Method:  http.MethodGet,
			Pattern: "/hello/:Name",
			APIFunc: s.HTTPHelloLab6,
		},
		{
			Name:    "Lab6 Hello POST",
			Method:  http.MethodPost,
			Pattern: "/hello",
			APIFunc: s.HTTPPostHelloLab6,
		},
	}
}
func (s *Server) HTTPPostHelloLab6(c *gin.Context) {
	var request Lab6HelloRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid JSON body",
		})
		return
	}

	s.Processor().PostHelloLab6(c, request.Name, request.Message)
}

func (s *Server) HTTPHelloLab6(c *gin.Context) {
	name := c.Param("Name")

	if name == "" {
		c.String(http.StatusBadRequest, "No name provided")
		return
	}
	s.Processor().HelloLab6(c, name)
}
