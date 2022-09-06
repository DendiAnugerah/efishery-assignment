package categorycontroller

import (
	"efishery-assignment/infrastructure/config"
	"efishery-assignment/model"
	"efishery-assignment/repository"
	categoryusecase "efishery-assignment/usecase/category"
	"strconv"

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

func (controller *CategoryController) GetById(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	categoryI, err := controller.Interact.GetById(id)

	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, categoryI)
	return
}

func (controller *CategoryController) UpdateCategory(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	NewCategory := model.Category{Id_category: id}
	NewCategory, err = controller.Interact.Update(NewCategory)

	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, NewCategory)
	return
}

func (controller *CategoryController) DeleteCategory(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	NewCategory := model.Category{Id_category: id}

	err = controller.Interact.Delete(NewCategory)
	if err != nil {
		c.JSON(500, err)
		return
	}

	type M map[string]interface{}
	c.JSON(200, M{"message": "Category deleted"})
	return
}
