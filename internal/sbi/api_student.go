package sbi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type StudentRequest struct {
	Name string `json:"name"`
}

func (s *Server) getStudentRoute() []Route {
	return []Route{
		{
			Name:    "Get Student",
			Method:  http.MethodGet,
			Pattern: "/:name",
			APIFunc: s.HTTPGetStudent,
		},
		{
			Name:    "Create Student",
			Method:  http.MethodPost,
			Pattern: "/",
			APIFunc: s.HTTPCreateStudent,
		},
	}
}

func (s *Server) HTTPGetStudent(c *gin.Context) {
	name := c.Param("name")

	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "student name is required",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"name":    name,
		"message": "Hello " + name,
	})
}

func (s *Server) HTTPCreateStudent(c *gin.Context) {
	var request StudentRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	if request.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "student name is required",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"name":    request.Name,
		"message": "Student created",
	})
}
