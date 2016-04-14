package api

import (
	"github.com/labstack/echo"
	"net/http"
)

type Sudoku struct {
	Grid string
	Size int
}

func CreateSudoku(c echo.Context) error {
	s := new(Sudoku)
	if err := c.Bind(s); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, s)
}