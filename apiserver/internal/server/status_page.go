package server

import (
	"github.com/gin-gonic/gin"
	"github.com/metoro-io/statusphere/common/api"
	"net/http"
	"strings"
)

type StatusPageResponse struct {
	StatusPage api.StatusPage `json:"statusPage"`
}

// StatusPage is a handler for the /status-page endpoint.
// It has a required query parameter of statusPageUrl XOR statusPageName
func (s *Server) statusPage(context *gin.Context) {
	statusPageUrl := context.Query("statusPageUrl")
	statusPageName := strings.ToLower(context.Query("statusPageName"))

	if statusPageUrl == "" && statusPageName == "" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "statusPageUrl or statusPageId is required"})
		return
	}

	if statusPageUrl != "" && statusPageName != "" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "statusPageUrl and statusPageId are mutually exclusive"})
		return
	}

	if statusPageUrl != "" {
		statusPage, found := s.statusPageCache.Get(statusPageUrl)
		if !found {
			context.JSON(http.StatusNotFound, gin.H{"error": "status page not known to statusphere"})
			return
		}
		context.JSON(http.StatusOK, StatusPageResponse{StatusPage: statusPage.(api.StatusPage)})
		return
	}

	if statusPageName != "" {
		for _, statusPage := range s.statusPageCache.Items() {
			if strings.ToLower(statusPage.Object.(api.StatusPage).Name) == statusPageName {
				context.JSON(http.StatusOK, StatusPageResponse{StatusPage: statusPage.Object.(api.StatusPage)})
				return
			}
		}
		context.JSON(http.StatusNotFound, gin.H{"error": "status page not known to statusphere"})
	}
}
