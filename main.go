package main

import (
	"log/slog"
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	meetRepo := NewMeetRepository()
	meetService := NewMeetService(meetRepo)
	NewMeetHandler(app, meetService)

	app.Logger.Fatal(app.Start(":4000"))
}

func Render(ctx echo.Context, statusCode int, t templ.Component) error {
	ctx.Response().Writer.WriteHeader(statusCode)
	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return t.Render(ctx.Request().Context(), ctx.Response().Writer)
}

type MeetHandler struct {
	service *MeetService
}

func NewMeetHandler(e *echo.Echo, service *MeetService) {
	handler := &MeetHandler{service: service}
	e.GET("/meets", handler.GetAllMeets)
	e.POST("/meets", handler.CreateMeet)
}

func (h *MeetHandler) GetAllMeets(c echo.Context) error {
	meets, err := h.service.GetMeets(c.Request().Context())

	if err != nil {
		echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return Render(c, http.StatusOK, Home(meets))
}

func (h *MeetHandler) CreateMeet(c echo.Context) error {
	meet := &Meet{}

	if err := c.Bind(meet); err != nil {
		slog.Error(err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid meet data")
	}

	err := h.service.CreateMeet(c.Request().Context(), meet)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Meet creation failed")
	}

	return c.JSON(http.StatusOK, meet)
}
