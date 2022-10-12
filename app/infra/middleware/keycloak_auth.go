package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/felixa1996/go_next_be/app/infra/iam"
	"github.com/labstack/echo/v4"
)

type keycloakAuth struct {
	handler     echo.HandlerFunc
	keycloakIam iam.KeycloakIAM
}

func KeycloakValidateJwt(keycloakIam iam.KeycloakIAM) echo.MiddlewareFunc {
	return func(h echo.HandlerFunc) echo.HandlerFunc {
		m := &keycloakAuth{
			handler:     h,
			keycloakIam: keycloakIam,
		}
		return m.handle
	}
}

func (m *keycloakAuth) handle(c echo.Context) error {
	req := c.Request()
	c.SetRequest(req)

	resp := c.Response()
	var handlerErr error
	defer func() {
		if v := recover(); v != nil {
			err, ok := v.(error)
			if !ok {
				err = errors.New(fmt.Sprint(v))
			}
			c.Error(err)
		}
	}()

	authHeader := c.Request().Header.Get("Authorization")
	if len(authHeader) < 1 {
		return echo.NewHTTPError(http.StatusUnauthorized, "Please provide authorization header")
	}

	if !strings.Contains(authHeader, "Bearer") {
		return echo.NewHTTPError(http.StatusUnauthorized, "Please provide bearer in authorization header")
	}

	accessToken := strings.Split(authHeader, " ")[1]

	rptResult, err := m.keycloakIam.RetrospectToken(c.Request().Context(), accessToken)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Please provide valid credentials")
	}
	isTokenValid := *rptResult.Active
	if !isTokenValid {
		return echo.NewHTTPError(http.StatusUnauthorized, "Token is expired")
	}

	handlerErr = m.handler(c)
	if handlerErr != nil {
		resp.Status = http.StatusInternalServerError
		if handlerErr, ok := handlerErr.(*echo.HTTPError); ok {
			resp.Status = handlerErr.Code
			reqPath := req.URL.RawPath
			if reqPath == "" {
				// nolint
				reqPath = req.URL.Path
			}
		}
	} else if !resp.Committed {
		resp.WriteHeader(http.StatusOK)
	}
	return handlerErr
}
