package server

import (
	"github.com/gin-gonic/gin"
	"github.com/metoro-io/statusphere/common/api"
	"net/http"
)

type StatusPagesResponse struct {
	StatusPages []api.StatusPage `json:"statusPages"`
}

func (s *Server) statusPages(context *gin.Context) {
	var statusPages []api.StatusPage
	for _, statusPage := range s.statusPageCache.Items() {
		statusPages = append(statusPages, statusPage.Object.(api.StatusPage))
	}

	context.JSON(http.StatusOK, StatusPagesResponse{StatusPages: statusPages})
}
