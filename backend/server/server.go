package server

import (
	"embed"
	"github.com/labstack/echo/v4"
	"github.com/zfz7/tetris_go/backend/helpers"
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

var serverPort = helpers.GetEnv("SERVER_PORT", "443")
var tetrisUrl = helpers.GetEnv("TETRIS_URL", "")
var fullChainPath =  "/etc/letsencrypt/live/"+tetrisUrl+"/fullchain.pem"
var privkeyPath = "/etc/letsencrypt/live/"+tetrisUrl+"/privkey.pem"
func StartWebServer(){
	e := echo.New()
	e.FileFS("/", "index.html", buildIndexHtml)
	e.StaticFS("/", buildDirFS)
	e.GET("/api/hello", helloWorldEndPoint)

	if _, err := os.Stat(fullChainPath); err == nil {
		e.Logger.Fatal(e.StartTLS(":"+ serverPort, fullChainPath, privkeyPath))
	} else {
		e.Logger.Fatal(e.Start(":"+ serverPort))
	}
}
