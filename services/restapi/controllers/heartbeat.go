package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/rintik-io/rintik-auth/internal/version"
)

// Heartbeat :
type Heartbeat struct{}

// Heartbeat godoc
// @Summary      Check Heartbeat of the Service
// @Description  return heartbeat status
// @Tags         heartbeat
// @Accept       json
// @Produce      json
// @Success      200  {object}  version.ResponseVersion
// @Router       /heartbeat [get]
func (p *Heartbeat) GetVersion(c *gin.Context) {
	// Set default response body
	responseStruct := version.GetVersion()

	// golibs.RespondNoStruct(w, http.StatusOK, &responseStruct)
	c.JSON(http.StatusOK, &responseStruct)
}
