package handler

import (
	"net/http"
	"sort"
	"strings"

	"github.com/SimonRininsland/gofino/model"
	"github.com/labstack/echo"
)

// SortString - Helper function to sort Strings
func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

// SortUser - User
func SortUser(c echo.Context) (err error) {
	/*
		// store names to later add and remove
			cookie := new(http.Cookie)
			cookie.Name = "username"
			cookie.Value = "jon"
			cookie.Expires = time.Now().Add(24 * time.Hour)
			c.SetCookie(cookie)
	*/

	u := new(model.AnomUser)
	if err = c.Bind(u); err != nil {
		return
	}

	//sorted := SortString(u)
	//return c.JSON(http.StatusOK, sorted)
	return c.JSON(http.StatusOK, u)
}
