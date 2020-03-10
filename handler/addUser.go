package handler

import (
	"net/http"

	"github.com/SimonRininsland/gofino/model"
	"github.com/labstack/echo"
)

// AddUser User
func AddUser(c echo.Context) (err error) {
	u := new(model.User)
	if err = c.Bind(u); err != nil {
		return
	}
	return c.JSON(http.StatusOK, u)
}
