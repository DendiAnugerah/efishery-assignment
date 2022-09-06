package categorycontroller

import (
	"efishery-assignment/infrastructure/config"
	"efishery-assignment/model"
	"efishery-assignment/repository"
	categoryusecase "efishery-assignment/usecase/category"

	"github.com/labstack/echo/v4"
)

type CategoryController struct {
	Interact categoryusecase.CateInteract
}

func NewCategoryController(dbhandler config.DbHandler) *CategoryController {
	return &CategoryController{
		Interact: categoryusecase.CateInteract{
			CatRepo: &repository.CatRepository{
				DbHandler: dbhandler,
			},
		},
	}
}

func (controller *CategoryController) CreateCategory(c echo.Context) (err error) {
	cat := model.Category{}
	c.Bind(&cat)

	categoryI, err := controller.Interact.Add(cat)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(201, categoryI)
	return
}

func (controller *CategoryController) GetAll(c echo.Context) (err error) {
	categoryI, err := controller.Interact.GetAll()
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, categoryI)
	return
}
