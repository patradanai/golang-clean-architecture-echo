package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Rest interface {
}

type MetaRest struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (r *MetaRest) String(message string) string {
	jsonData, _ := json.Marshal(r.Data)
	return fmt.Sprintf("%v%v", r.Message, jsonData)
}

type MetaRestPagination struct {
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	TotalCount int         `json:"total_count"`
}

func (r *MetaRestPagination) String(message string) string {
	jsonData, _ := json.Marshal(r.Data)
	return fmt.Sprintf("%v%v", r.Message, jsonData)
}

func Response[T any](c echo.Context, message string, data T) error {
	return c.JSON(http.StatusOK, &MetaRest{
		Message: message,
		Data:    data,
	})
}

func ResponsePagination[T any](c echo.Context, message string, data T, count int) error {
	return c.JSON(http.StatusOK, &MetaRestPagination{
		Message:    message,
		Data:       data,
		TotalCount: count,
	})
}
