package getme

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

var usecase *Usecase

func HTTPv1(c echo.Context) error {
	var input Input
	binder := new(echo.DefaultBinder)
	if err := binder.BindHeaders(c, &input); err != nil {
		logrus.Debugf("debug me: %s", err)
		return err
	}

	logrus.Debug("token: %w", input)
	logrus.Debug(input.Token)

	ctx := c.Request().Context()

	output, err := usecase.GetUserByToken(ctx, input)
	if err != nil {
		return err
	}

	return c.JSON(200, output)
}
