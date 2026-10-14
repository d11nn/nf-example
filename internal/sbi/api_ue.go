package sbi

import (
	"net/http"
	"strings"

	nf_context "github.com/Alonza0314/nf-example/internal/context"
	"github.com/Alonza0314/nf-example/internal/logger"
	"github.com/gin-gonic/gin"
)

func (s *Server) getUeRoute() []Route {
	return []Route{
		{
			Name:    "List UEs",
			Method:  http.MethodGet,
			Pattern: "/",
			APIFunc: s.HTTPListUes,
			// Use
			// curl -X GET http://127.0.0.163:8000/ue/ -w "\n"
		},
		{
			Name:    "Get UE",
			Method:  http.MethodGet,
			Pattern: "/:Supi",
			APIFunc: s.HTTPGetUe,
			// Use
			// curl -X GET http://127.0.0.163:8000/ue/imsi-208930000000001 -w "\n"
		},
		{
			Name:    "Register UE",
			Method:  http.MethodPost,
			Pattern: "/",
			APIFunc: s.HTTPRegisterUe,
			// Use
			// curl -X POST http://127.0.0.163:8000/ue/ -H "Content-Type: application/json" \
			//   -d '{"supi":"imsi-208930000000001","gnbId":"gnb-1","dnn":"internet"}' -w "\n"
		},
	}
}

func (s *Server) HTTPListUes(c *gin.Context) {
	logger.SBILog.Infof("In HTTPListUes")

	s.Processor().ListUes(c)
}

func (s *Server) HTTPGetUe(c *gin.Context) {
	logger.SBILog.Infof("In HTTPGetUe")

	supi := c.Param("Supi")
	if supi == "" {
		c.String(http.StatusBadRequest, "No supi provided")
		return
	}

	s.Processor().GetUe(c, supi)
}

func (s *Server) HTTPRegisterUe(c *gin.Context) {
	logger.SBILog.Infof("In HTTPRegisterUe")

	var ue nf_context.UeInfo
	if err := c.ShouldBindJSON(&ue); err != nil {
		c.String(http.StatusBadRequest, "Invalid JSON body")
		return
	}
	if !strings.HasPrefix(ue.Supi, "imsi-") {
		c.String(http.StatusBadRequest, "supi must start with imsi-")
		return
	}

	s.Processor().RegisterUe(c, ue)
}
