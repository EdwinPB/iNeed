package app

import (
	"net/http"

	"ineedApp/public"
	"ineedApp/app/actions/home"
	"ineedApp/app/middleware"

	"github.com/gobuffalo/buffalo"
)

// SetRoutes for the application
func setRoutes(root *buffalo.App) {
	root.Use(middleware.RequestID)
	root.Use(middleware.Database)
	root.Use(middleware.ParameterLogger)
	root.Use(middleware.CSRF)

	root.GET("/", home.Index)
	root.ServeFiles("/", http.FS(public.FS()))
}