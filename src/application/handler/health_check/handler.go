package healthcheck

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthCheckHandler struct{}

func NewHealthCheckHandler() *HealthCheckHandler {
	return &HealthCheckHandler{}
}

func (c *HealthCheckHandler) Run(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, nil)
}
