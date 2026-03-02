package httpcontroller

import (
	"backend/internal/chat/get_chat_messages"
	getchats "backend/internal/chat/get_chats"
	"backend/internal/user/login_user"
	"backend/internal/user/register_user"
	"backend/internal/user/search_user"
	"backend/internal/wsserver"
	jwttoken "backend/pkg/jwt_token"
	"backend/pkg/metrics"
	pkg_middleware "backend/pkg/middleware"
	"embed"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

//go:embed docs
var docsFS embed.FS

func Router(ws *wsserver.WsServer, jwtManager *jwttoken.JWTManager) http.Handler {

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173", "http://localhost:8004", "http://xn----7sbbozvcgr0a7c4b.xn--p1ai:5173", "http://xn----7sbbozvcgr0a7c4b.xn--p1ai:8004"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
	}))

	e.Use(middleware.Recover())
	e.Use(middleware.RequestLogger())
	e.Use(metrics.RequestCounter())

	e.GET("/ws", func(c echo.Context) error {
		ws.Handler(c.Response(), c.Request())
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

	v1.GET("/metrics", echo.WrapHandler(promhttp.Handler()))

	protected := v1.Group("", pkg_middleware.AuthMiddleware(jwtManager))
	protected.GET("/users", search_user.HTTPv1)
	protected.GET("/chats", getchats.HTTPv1)
	protected.GET("/chats/:id/messages", get_chat_messages.HTTPv1)

	return e
}
