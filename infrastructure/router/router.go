package router

import (
	categorycontroller "efishery-assignment/controller/category"
	"efishery-assignment/infrastructure/config"

	"github.com/labstack/echo/v4"
)

func Init() {
	// initialize echo framework
	e := echo.New()

	// -- Category --
	category := e.Group("category")
	custControl := categorycontroller.NewCategoryController(config.Conection())

	category.POST("/", func(c echo.Context) error { return custControl.CreateCategory(c) })
	category.GET("/", func(c echo.Context) error { return custControl.GetAll(c) })
	category.GET("/:id", func(c echo.Context) error { return custControl.GetById(c) })
	category.PUT("/:id", func(c echo.Context) error { return custControl.UpdateCategory(c) })
	category.DELETE("/:id", func(c echo.Context) error { return custControl.DeleteCategory(c) })

	// -- Product --

	// // start service secho
	e.Logger.Fatal(e.Start(":9090"))
}
