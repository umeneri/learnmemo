package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthController interface {
	Index(c *gin.Context)
}

type healthController struct {
}

func NewHealthController() HealthController {
	return &healthController{}
}

func (t *healthController) Index(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}