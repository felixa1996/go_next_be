package healthcheck

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type livenessProbes struct {
	Message string `json:"message"`
}

type readinessProbes struct {
	Key     string `json:"key"`
	IsReady bool   `json:"is_ready"`
}

func RegisterHealthCheckHandler(logger *zap.Logger, group *echo.Group) {
	group.GET("/liveness", livenessProbeHandler)
	group.GET("/readiness", readinessProbeHandler)
}

// Probes godoc
// @Summary      Liveness probe
// @Description  Liveness probe
// @Tags         Probe
// @Produce      json
// @Success      200  {object}  healthcheck.livenessProbes{message=string}
// @Router       /probes/liveness [get]
func livenessProbeHandler(c echo.Context) error {
	msg := &livenessProbes{
		Message: "Application ready to serve",
	}
	return c.JSON(http.StatusOK, msg)
}

// Probes godoc
// @Summary      Readiness probe
// @Description  Readiness probe
// @Tags         Probe
// @Produce      json
// @Success      200  {object}  []healthcheck.readinessProbes{key=string,is_ready=bool}
// @Router       /probes/readiness [get]
func readinessProbeHandler(c echo.Context) error {
	msg := []readinessProbes{
		{Key: "Mongo", IsReady: GetMongoReadiness()},
		{Key: "Aws Session", IsReady: GetAwsSessionReadiness()},
	}
	return c.JSON(http.StatusOK, msg)
}
