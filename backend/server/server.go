package server

import (
	"embed"
	"github.com/labstack/echo/v4"
	"github.com/zfz7/tetris_go/backend/helpers"
)

var (
	//go:embed build
	dist embed.FS
	//go:embed build/index.html
	indexHTML      embed.FS
	buildDirFS     = echo.MustSubFS(dist, "build")
	buildIndexHtml = echo.MustSubFS(indexHTML, "build")
)

var serverPort = helpers.GetEnv("SERVER_PORT", "443")

func StartWebServer(){
	e := echo.New()
	e.FileFS("/", "index.html", buildIndexHtml)
	e.StaticFS("/", buildDirFS)
	e.GET("/api/hello", helloWorldEndPoint)
	e.Logger.Fatal(e.Start(":"+ serverPort))
}
