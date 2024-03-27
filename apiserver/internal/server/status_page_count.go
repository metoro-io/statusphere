package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type StatusPageCountResponse struct {
	StatusPageCount int `json:"statusPageCount"`
}

// statusPageCount is a handler for the /statusPages/count endpoint.
func (s *Server) statusPageCount(context *gin.Context) {
	context.JSON(http.StatusOK, StatusPageCountResponse{StatusPageCount: s.statusPageCache.ItemCount()})
}
