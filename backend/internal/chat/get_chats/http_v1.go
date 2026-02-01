package getchats

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	pkg_middleware "backend/internal/middleware"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

var usecase *Usecase

func HTTPv1(c echo.Context) error {
	var input Input
	input, err := parseQueryParamsWithInput(c, input)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid query params: %v", err))
	}

	claims, ok := pkg_middleware.GetUserClaims(c.Request().Context())
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, "claims not found")
	}
	input.UserID = claims.UserID

	logrus.Infof("input: %v", input)

	output, err := usecase.GetChats(context.Background(), input)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, output)
}

func parseQueryParamsWithInput(c echo.Context, input Input) (Input, error) {
	limitMessages := c.QueryParam("limit_messages")
	if limitMessages == "" {
		limitMessages = "5"
	}

	limitMessagesInt, err := strconv.Atoi(limitMessages)
	if err != nil {
		return input, err
	}
	input.LimitMessages = limitMessagesInt
	return input, nil
}
