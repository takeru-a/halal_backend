package usecase

import (
	"github.com/gin-gonic/gin"
	"github.com/takeru-a/halal_backend/lib"
)

type DeleteShopUsecase struct {
	shopRepo lib.ShopRepository
}

func NewDeleteShopUsecase(
	shopRepo lib.ShopRepository,
) *DeleteShopUsecase {
	return &DeleteShopUsecase{
		shopRepo: shopRepo,
	}
}

func (usecase *DeleteShopUsecase) Execute(
	ctx *gin.Context,
	id string,
) error {
	_, err := usecase.shopRepo.FindById(ctx, id)
	if err != nil {
		return err
	}
	err = usecase.shopRepo.DeleteById(ctx, id)
	if err != nil {
		return err
	}
	return nil
}