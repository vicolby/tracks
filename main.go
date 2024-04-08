package main

import (
	"net/http"
	"strconv"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

var testData = TrackData{
	Tracks: []Track{
		{
            ID: 0,
			Name:            "Learning",
			MaxProgress:     100, // Chapters
			CurrentProgress: 50,  // Chapters
		},
		{
            ID: 1,
			Name:            "Eating",
			MaxProgress:     100, // Meals
			CurrentProgress: 20,  // Meals
		},
		{
            ID: 2,
			Name:            "Sleeping",
			MaxProgress:     100, // Hours
			CurrentProgress: 60,  // Hours
		},
	},
}

func main() {
	app := echo.New()
	app.GET("/", HomeHandler)
	app.POST("/track/:id", IncrementTrack)
	app.Logger.Fatal(app.Start(":4000"))
}

func Render(ctx echo.Context, statusCode int, t templ.Component) error {
	ctx.Response().Writer.WriteHeader(statusCode)
	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return t.Render(ctx.Request().Context(), ctx.Response().Writer)
}

func HomeHandler(c echo.Context) error {
	return Render(c, http.StatusOK, Home(testData))
}

func IncrementTrack(c echo.Context) error {
	trackId := c.Param("id")
	track, _ := strconv.Atoi(trackId)
	testData.Tracks[track].CurrentProgress += 1
	return Render(c, http.StatusOK, TrackProgress(testData.Tracks[track]))
}
