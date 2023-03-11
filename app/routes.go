package app

import (
	"net/http"

	"ineedApp/app/actions"
	"ineedApp/app/actions/home"
	"ineedApp/app/middleware"
	"ineedApp/public"

	"github.com/gobuffalo/buffalo"
)

// SetRoutes for the application
func setRoutes(root *buffalo.App) {
	root.Use(middleware.RequestID)
	root.Use(middleware.Database)
	root.Use(middleware.ParameterLogger)
	root.Use(middleware.CSRF)

	root.GET("/", home.Index)

	// API to access business info
	business := actions.BusinessesResource{}
	businessGroup := root.Group("/business")
	businessGroup.GET("/", business.ListBussines)
	businessGroup.GET("/list", business.List)

	root.ServeFiles("/", http.FS(public.FS()))
}
