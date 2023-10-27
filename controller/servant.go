package controller

import (
    "github.com/labstack/echo/v4"
    "net/http"
)

func SearchServant(c echo.Context) (err error) {
    query := c.QueryParam("query")
    return c.String(http.StatusOK, query)
}

func GetServant(c echo.Context) (err error) {
    id := c.Param("id")
    return c.String(http.StatusOK, id)
}
