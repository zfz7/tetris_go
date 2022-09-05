package main

import (
	"embed"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
)

var (
	//go:embed build
	dist embed.FS
	//go:embed build/index.html
	indexHTML      embed.FS
	buildDirFS     = echo.MustSubFS(dist, "build")
	buildIndexHtml = echo.MustSubFS(indexHTML, "build")
)

// TODO save this in DB
var visitorCount = 0
var ServerPort = getEnv("SERVER_PORT", "443")
func main() {
	e := echo.New()
	e.FileFS("/", "index.html", buildIndexHtml)
	e.StaticFS("/", buildDirFS)
	e.GET("/api/hello", helloWorldEndPoint)
	e.Logger.Fatal(e.Start(":"+ServerPort))
}
type HelloDTO struct {
	Message  string `json:"message" xml:"message"`
}
func helloWorldEndPoint(context echo.Context) error {
	dto := &HelloDTO{
		Message:  fmt.Sprint("You are visitor ", visitorCount),
	}
	visitorCount++
	return context.JSON(http.StatusOK, dto)
}
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}