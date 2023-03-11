package usecase

import (
	"github.com/gin-gonic/gin"
	"github.com/takeru-a/halal_backend/lib"
	"github.com/takeru-a/halal_backend/models"
)

type GetShopUsecase struct {
	shopRepo lib.ShopRepository
}

func NewGetShopUsecase(
	shopRepo lib.ShopRepository,
) *GetShopUsecase {
	return &GetShopUsecase{
		shopRepo: shopRepo,
	}
}

func (usecase *GetShopUsecase) Execute(
	ctx *gin.Context,
	id string,
) (*models.Shop, error) {
	shop, err := usecase.shopRepo.FindById(ctx, id)
	if err != nil {
		return nil, err
	}
	return shop, nil
}