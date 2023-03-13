package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/takeru-a/halal_backend/controllers"
	"github.com/takeru-a/halal_backend/lib"
	"github.com/takeru-a/halal_backend/usecase"
)
var connection string = fmt.Sprintf("host=halalDB user=%s password=%s dbname=%s sslmode=disable",
os.Getenv("POSTGRES_USER"), 
os.Getenv("POSTGRES_PASSWORD"), 
os.Getenv("POSTGRES_DB"))
func generateDB() (*sql.DB, error) {
	return sql.Open(
		"postgres", connection)
}

func main(){
	router := gin.Default()
	db, dbErr := generateDB()
	if dbErr != nil {
		panic("failed database connection")
	}
	err := db.Ping()
	if err !=nil{
		log.Fatal(err)
	}
	
	shopRepo := lib.NewShopRepository(db)
	createShopUsecase := usecase.NewCreateShopUsecase(shopRepo)
	getShopUsecase := usecase.NewGetShopUsecase(shopRepo)
	listShopUsecase := usecase.NewListShopUsecase(shopRepo)
	updateShopUsecase := usecase.NewUpdateShopUsecase(shopRepo)
	deleteShopUsecase := usecase.NewDeleteShopUsecase(shopRepo)
	
	shopCon := controllers.NewShopController(
		*createShopUsecase,
		*listShopUsecase,
		*getShopUsecase,
		*updateShopUsecase,
		*deleteShopUsecase,
	)
	router.GET("/", func(ctx *gin.Context){
		ctx.JSON(200, gin.H{"msg":"hello!,World!"})
	} )
	router.GET("/shops", shopCon.ListShops)
	router.GET("/shops/:id", shopCon.GetShopByID)
	router.POST("/shops", shopCon.CreateShop)
	router.PUT("/shops", shopCon.UpdateShop)
	router.DELETE("/shops/:id", shopCon.DeleteShop)
	router.Run()
}