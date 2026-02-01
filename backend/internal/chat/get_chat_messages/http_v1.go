package get_chat_messages

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
	input, err := parseInput(c, input)
	if err != nil {
		return err
	}

	claims, ok := pkg_middleware.GetUserClaims(c.Request().Context())
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, "claims not found")
	}
	input.UserID = claims.UserID

	//TODO: сделать проверку на пренадлежность пользователя к чату

	logrus.Infof("input: %v", input)

	output, err := usecase.GetChatMessages(context.Background(), input)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, output)
}

func parseInput(c echo.Context, input Input) (Input, error) {
	id := c.Param("id")
	if id == "" {
		return input, echo.NewHTTPError(http.StatusBadRequest, "chat_id is required")
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return input, echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid chat_id: %v", err))
	}
	input.ChatID = idInt

	limitStr := c.QueryParam("limit")
	if limitStr == "" {
		return input, echo.NewHTTPError(http.StatusBadRequest, "limit is required")
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		return input, echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid limit: %v", err))
	}
	input.Limit = limit

	offsetStr := c.QueryParam("offset")
	if offsetStr == "" {
		offsetStr = "0"
	}
	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		return input, echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid offset: %v", err))
	}
	input.Offset = offset

	return input, nil
}
