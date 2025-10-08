package sbi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) getAnyaRoute() []Route {
	return []Route{
		{
			Name:    "GetAnya",
			Method:  http.MethodGet,
			Pattern: "/:name",
			APIFunc: s.HTTPGetAnya,
		},
		{
			Name:    "CreateAnya",
			Method:  http.MethodPost,
			Pattern: "/",
			APIFunc: s.HTTPCreateAnya,
		},
	}
}

func (s *Server) HTTPGetAnya(c *gin.Context) {
	name := c.Param("name")
	s.Processor().FindAnyaCharacterName(c, name)
}

func (s *Server) HTTPCreateAnya(c *gin.Context) {
	s.Processor().CreateAnyaCharacter(c)
}
