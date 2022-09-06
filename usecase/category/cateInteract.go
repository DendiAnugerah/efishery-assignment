package categoryusecase

import "efishery-assignment/model"

type CateInteract struct {
	CatRepo CatRepo
}

func (i *CateInteract) Add(mc model.Category) (c model.Category, err error) {
	c, err = i.CatRepo.Create(mc)
	return
}

func (i *CateInteract) GetById(id int) (c model.Category, err error) {
	c, err = i.CatRepo.GetById(id)
	return
}

func (i *CateInteract) GetAll() (c model.Categories, err error) {
	c, err = i.CatRepo.GetAll()
	return
}

func (i *CateInteract) Update(mc model.Category) (c model.Category, err error) {
	c, err = i.CatRepo.Update(mc)
	return
}

func (i *CateInteract) Delete(c model.Category) (err error) {
	err = i.CatRepo.Delete(c)
	return
}
