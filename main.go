package main

import (
	"github.com/labstack/echo/v4"
	"urlcutter/app"
	"urlcutter/infrastructure/httphandlers"
	"urlcutter/infrastructure/localservices"
	repo2 "urlcutter/infrastructure/repo"
)

func main() {
	e := echo.New()

	repo := repo2.NewKeysaver()

	geyGeneratorService := localservices.NewKeyGenerator()

	service := app.NewUrlcutterSrvice(*geyGeneratorService, repo)

	handler := httphandlers.NewUrlGenerator(service)

	e.POST("/urlcutter", handler.UrlCutter)
	e.GET("/:key", handler.GetUrl)
	e.Logger.Fatal(e.Start(":8088"))
}
