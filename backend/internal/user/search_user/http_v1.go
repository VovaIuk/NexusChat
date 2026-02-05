package search_user

import (
	"context"
	"strconv"

	"github.com/labstack/echo/v4"
)

var usecase *Usecase

func HTTPv1(c echo.Context) error {
	var input Input
	input, err := parseQueryParamsForInput(c, input)
	if err != nil {
		return err
	}

	output, err := usecase.SearchUser(context.Background(), input)
	if err != nil {
		return err
	}

	return c.JSON(200, output)

}

func parseQueryParamsForInput(c echo.Context, input Input) (Input, error) {
	limit := c.QueryParam("limit")
	if limit == "" {
		limit = "5"
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		return input, echo.NewHTTPError(400, "invalid limit: must be a number")
	}
	input.Limit = limitInt

	tag := c.QueryParam("tag")
	if tag == "" {
		return input, echo.NewHTTPError(400, "tag is required")
	}
	input.Tag = tag
	return input, nil
}
