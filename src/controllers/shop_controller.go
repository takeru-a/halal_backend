package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/takeru-a/halal_backend/models/request"
	"github.com/takeru-a/halal_backend/usecase"
)

type ShopController struct {
	createShopUsecase usecase.CreateShopUsecase
	listShopUsecase   usecase.ListShopUsecase
	getShopUsecase    usecase.GetShopUsecase
	
}

func NewShopController(
	createShopUsecase usecase.CreateShopUsecase,
	listShopUsecase   usecase.ListShopUsecase,
	getShopUsecase    usecase.GetShopUsecase,
) *ShopController {
	return &ShopController{
	createShopUsecase : createShopUsecase,
	listShopUsecase   : listShopUsecase,
	getShopUsecase    : getShopUsecase,
	}
}

func (con *ShopController) ListShops(ctx *gin.Context) {
	shops, err := con.listShopUsecase.Execute(ctx)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	ctx.IndentedJSON(http.StatusOK, shops)
}


func (con *ShopController) GetShopByID(ctx *gin.Context) {
	id := ctx.Param("id")
	shop, err := con.getShopUsecase.Execute(ctx, id)

	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, shop)
}


func (con *ShopController) CreateShop(ctx *gin.Context) {
	var newShop request.ShopCreateRequest
	if err := ctx.BindJSON(&newShop); err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed bind json"})
		return
	}
	err := con.createShopUsecase.Execute(
		ctx,
		newShop.ID,
		newShop.Name,
		newShop.Introduction,
		newShop.Location,
		newShop.Level,
		newShop.Score,
	)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, "OK")
}
