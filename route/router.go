package route

import (
	"github.com/Reyurnible/Sudoku-server/api"
	"github.com/labstack/echo"
	middleware "github.com/labstack/echo/middleware"
)

func Init() *echo.Echo {
	e := echo.New()
	// Debug
	e.Debug()
	// Setup Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Gzip())
	v1 := e.Group("/api/v1")
	{
		v1.Post("/sudoku", api.CreateSudoku)
	}
	return e
}
