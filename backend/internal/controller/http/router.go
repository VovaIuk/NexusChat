package httpcontroller

import (
	"backend/internal/chat/get_chatheaders"
	"backend/internal/chat/get_chathistory"
	pkg_middleware "backend/internal/middleware"
	"backend/internal/user/login_user"
	"backend/internal/user/register_user"
	"backend/internal/wsserver"
	"embed"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//TODO: переписать на Echo

//go:embed docs
var docsFS embed.FS

func Router(ws *wsserver.WsServer) http.Handler {

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173", "http://localhost:8004"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
	}))

	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	e.GET("/ws", func(c echo.Context) error {
		ws.WsHandler(c.Response(), c.Request())
		return nil
	})

	e.StaticFS("/api/docs", echo.MustSubFS(docsFS, "docs"))

	api := e.Group("/api")
	v1 := api.Group("/v1")

	v1.POST("/login", func(c echo.Context) error {
		login_user.HTTP_V1(c.Response(), c.Request())
		return nil
	})

	v1.POST("/registration", func(c echo.Context) error {
		register_user.HTTP_V1(c.Response(), c.Request())
		return nil
	})

	private := v1.Group("/private", pkg_middleware.AuthMiddleware())

	private.GET("/chats/:id/history", func(c echo.Context) error {
		get_chathistory.HTTP_V1(c.Response(), c.Request())
		return nil
	})
	private.GET("/chats/headers", func(c echo.Context) error {
		get_chatheaders.HTTP_V1(c.Response(), c.Request())
		return nil
	})

	return e
}
