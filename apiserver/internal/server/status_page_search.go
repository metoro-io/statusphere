package server

import (
	"github.com/gin-gonic/gin"
	"github.com/lithammer/fuzzysearch/fuzzy"
	"github.com/metoro-io/statusphere/common/api"
	"math"
	"net/http"
	"sort"
)

type StatusPageSearchResponse struct {
	StatusPages []api.StatusPage `json:"statusPages"`
}

type statusPageRanked struct {
	StatusPage api.StatusPage `json:"statusPage"`
	Score      int            `json:"score"`
}

// statusPageSearch is a handler for the /statusPages/search endpoint.
// It has a required query parameter of query
// Entries returned first are the ones with the best match
// Limited to 25 results
func (s *Server) statusPageSearch(context *gin.Context) {
	query := context.Query("query")
	if query == "" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "query is required"})
		return
	}

	var statusPagesRanked []statusPageRanked

	for _, statusPage := range s.statusPageCache.Items() {
		score := math.MaxInt
		nameMatch := fuzzy.RankMatch(query, statusPage.Object.(api.StatusPage).Name)
		urlMatch := fuzzy.RankMatch(query, statusPage.Object.(api.StatusPage).URL)
		if nameMatch != -1 {
			score = nameMatch
		}
		if urlMatch != -1 {
			if urlMatch < score {
				score = urlMatch
			}
		}

		if score != math.MaxInt {
			statusPagesRanked = append(statusPagesRanked, statusPageRanked{StatusPage: statusPage.Object.(api.StatusPage), Score: score})
		}
	}

	// Sort the status pages by score
	sort.Slice(statusPagesRanked, func(i, j int) bool {
		return statusPagesRanked[i].Score < statusPagesRanked[j].Score
	})

	var statusPages []api.StatusPage
	for _, statusPage := range statusPagesRanked {
		statusPages = append(statusPages, statusPage.StatusPage)
	}

	if len(statusPages) > 25 {
		statusPages = statusPages[:25]
	}

	context.JSON(http.StatusOK, StatusPageSearchResponse{StatusPages: statusPages})
}
