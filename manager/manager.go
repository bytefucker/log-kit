package manager

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/yihongzhi/log-kit/config"
	"net/http"
)

type WebServer struct {
	port   string
	server *echo.Echo
}

func NewManagerServer(config *config.AppConfig) (*WebServer, error) {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/", hello)
	return &WebServer{
		server: e,
		port:   config.Port,
	}, nil
}

func (s *WebServer) Start() error {
	return s.server.Start(":" + s.port)
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
