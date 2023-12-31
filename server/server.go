package server

import (
	"github.com/golang-jwt/jwt"
	"github.com/jbillote/fgo-planner-api/controller"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"os"
)

type FGOPlannerAPI struct {
	e *echo.Echo
}

func NewServer() *FGOPlannerAPI {
	return &FGOPlannerAPI{
		e: echo.New(),
	}
}

func (s *FGOPlannerAPI) Start(port string) {
	s.e.Use(middleware.Logger())
	s.e.Use(middleware.Recover())
	s.e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST},
	}))
	s.e.Use(echojwt.JWT([]byte(os.Getenv("SIGNING_KEY"))))

	s.e.POST("/token", func(c echo.Context) error {
		token := jwt.New(jwt.SigningMethodHS256)
		t, err := token.SignedString([]byte(os.Getenv("SIGNING_KEY")))
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, echo.Map{
			"token": t,
		})
	})

	fgo := s.e.Group("/api/fgo")
	v1 := fgo.Group("/v1")
	v1Routes(v1)

	s.e.Logger.Fatal(s.e.Start(port))
}

func v1Routes(e *echo.Group) {
	servantGroup := e.Group("/servant")
	servantGroup.GET("", controller.SearchServant)
	servantGroup.GET("/:id", controller.GetServant)
}

func (s *FGOPlannerAPI) Close() {
	s.e.Close()
}
