package usecase

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/takeru-a/halal_backend/lib"
	"github.com/takeru-a/halal_backend/models"
)

type CreateShopUsecase struct {
	shopRepo lib.ShopRepository
}

func NewCreateShopUsecase(
	shopRepo lib.ShopRepository,
) *CreateShopUsecase {
	return &CreateShopUsecase{
		shopRepo: shopRepo,
	}
}

func (usecase *CreateShopUsecase) Execute(
	ctx *gin.Context,
	id string,
	name string,
	introduction string,
	location string,
	level int,
	score float64,
) error {
	exist, err := usecase.shopRepo.FindById(ctx, id)
	if exist != nil {
		return errors.New("That shop already exists.")
	}
	newShop ,newShoErr := models.NewShop(id, name, introduction, location, level, score )
	if newShoErr != nil {
		return newShoErr
	}
	err = usecase.shopRepo.Save(ctx, *newShop)
	if err != nil {
		return err
	}
	return nil
}