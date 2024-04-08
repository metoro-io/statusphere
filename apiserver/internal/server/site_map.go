package server

import (
	"github.com/gin-gonic/gin"
	"github.com/ikeikeikeike/go-sitemap-generator/v2/stm"
	"github.com/metoro-io/statusphere/common/api"
	"net/http"
	"net/url"
)

func (s *Server) siteMap(context *gin.Context) {
	sm := stm.NewSitemap(1)
	sm.Create()
	pages := s.statusPageCache.Items()
	if len(pages) == 0 {
		s.logger.Warn("no status pages found")
		context.JSON(http.StatusInternalServerError, "no status pages found")
		return
	}
	for _, statusPage := range pages {
		page, ok := statusPage.Object.(api.StatusPage)
		if !ok {
			s.logger.Error("failed to cast status page")
			context.JSON(http.StatusInternalServerError, "failed to cast status page")
			return
		}
		escapeString := url.QueryEscape(page.Name)
		sm.Add(stm.URL{{"loc", "https://metoro.io/statusphere/status/" + escapeString}, {"changefreq", "always"}, {"mobile", true}, {"priority", 0.1}})
	}

	content := sm.XMLContent()
	context.Data(http.StatusOK, "application/xml", []byte(content))
}
