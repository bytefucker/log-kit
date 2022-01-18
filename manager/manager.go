package manager

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/yihongzhi/log-kit/config"
	"net/http"
)

type Server struct {
	port int
	echo *echo.Echo
}

func NewManagerServer(config *config.AppConfig) (*Server, error) {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/", hello)
	return &Server{
		echo: e,
		port: config.Port,
	}, nil
}

func (s *Server) Start() error {
	return s.echo.Start(fmt.Sprintf(":%d", s.port))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
