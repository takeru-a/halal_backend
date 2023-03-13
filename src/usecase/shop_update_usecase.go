package usecase

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/takeru-a/halal_backend/lib"
	// "github.com/takeru-a/halal_backend/models"
)

type UpdateShopUsecase struct {
	shopRepo lib.ShopRepository
}

func NewUpdateShopUsecase(
	shopRepo lib.ShopRepository,
) *UpdateShopUsecase {
	return &UpdateShopUsecase{
		shopRepo: shopRepo,
	}
}

func (usecase *UpdateShopUsecase) Execute(
	ctx *gin.Context,
	id string,
	name string,
	introduction string,
	location string,
	level int,
	score float64,
) error {
	exist, _ := usecase.shopRepo.FindById(ctx, id)
	if exist == nil {
		return errors.New("That shop doesn't exist.")
	}
	
	err := exist.Update(name, introduction, location, level, score )
	if err != nil {
		return err
	}
	err = usecase.shopRepo.Update(ctx, *exist)
	if err != nil {
		return err
	}
	return nil
}