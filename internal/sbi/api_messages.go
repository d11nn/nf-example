package sbi

import (
	"net/http"

	"github.com/Alonza0314/nf-example/internal/logger"
	"github.com/gin-gonic/gin"
)


type CreateMessageRequest struct {
    Name string `json:"name" binding:"required"`
    Food string `json:"food" binding:"required"`
}



func (s *Server) getmessageRoute() []Route {
	return []Route{
		{
			Name:    "Hello message!",
			Method:  http.MethodGet,
			Pattern: "/",
			APIFunc: func(c *gin.Context) {
				c.JSON(http.StatusOK, "Hello message!")
			},
			// Use
			// curl -X GET http://127.0.0.163:8000/spyfamily/ -w "\n"
		},
		{
			Name:    "message text",
			Method:  http.MethodGet,
			Pattern: "/:name",
			APIFunc: s.HTTPSerchmessagetext,
			// Use
			// curl -X GET http://127.0.0.163:8000/spyfamily/Anya -w "\n"
			// "Character: Anya Forger"
		},

		{
			Name:    "post message",
			Method:  http.MethodPost,
			Pattern: "/",
			APIFunc: s.HTTPinsert,
			// Use
			// curl -X GET http://127.0.0.163:8000/spyfamily/Anya -w "\n"
			// "Character: Anya Forger"
		},
	}
}

func (s *Server) HTTPSerchmessagetext(c *gin.Context) {
	logger.SBILog.Infof("In HTTPSerchCharacter")

	targetname := c.Param("name")
	if targetname == "" {
		c.String(http.StatusBadRequest, "No name provided")
		return
	}

	s.Processor().Findmessagetext(c, targetname)
}


func (s *Server) HTTPinsert(c *gin.Context) {
	logger.SBILog.Infof("In HTTPinsert")

	var req CreateMessageRequest
	if err:=c.ShouldBindJSON(&req) ; err !=nil{
		 c.JSON(http.StatusBadRequest, gin.H{
            "error": "請提供有效的 JSON，name 和 food 皆不可為空",
        })
        return
	}

	name :=req.Name
	food :=req.Food
	s.Processor().Insert(c,name,food)
}

