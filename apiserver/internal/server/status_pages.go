package server

import (
	"github.com/gin-gonic/gin"
	"github.com/metoro-io/statusphere/common/api"
	"net/http"
	"sort"
)

type StatusPagesResponse struct {
	StatusPages []api.StatusPage `json:"statusPages"`
}

func (s *Server) statusPages(context *gin.Context) {
	var statusPages []api.StatusPage
	for _, statusPage := range s.statusPageCache.Items() {
		statusPages = append(statusPages, statusPage.Object.(api.StatusPage))
	}

	// Sort the status pages by name alphabetically a to z
	sort.Slice(statusPages, func(i, j int) bool {
		return statusPages[i].Name < statusPages[j].Name
	})

	context.JSON(http.StatusOK, StatusPagesResponse{StatusPages: statusPages})
}
