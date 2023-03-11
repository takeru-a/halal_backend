package usecase

import (
	"github.com/gin-gonic/gin"
	"github.com/takeru-a/halal_backend/lib"
	"github.com/takeru-a/halal_backend/models"
)

type ListShopUsecase struct {
	shopRepo lib.ShopRepository
}

func NewListShopUsecase(
	shopRepo lib.ShopRepository,
) *ListShopUsecase {
	return &ListShopUsecase{
		shopRepo: shopRepo,
	}
}

func (usecase *ListShopUsecase) Execute(
	ctx *gin.Context,
) (*[]models.Shop, error) {
	shops, err := usecase.shopRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return shops, nil
}