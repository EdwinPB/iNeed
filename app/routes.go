package app

import (
	"fmt"
	"io"
	"net/http"

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
	root.POST("/test/", func(c buffalo.Context) error {
		b, err := io.ReadAll(c.Request().Body)
		if err != nil {
			c.Logger().Errorf("error getting body %v", err)
		}

		fmt.Println("---->", string(b))

		return c.Render(http.StatusOK, nil)
	})
	root.ServeFiles("/", http.FS(public.FS()))
}
