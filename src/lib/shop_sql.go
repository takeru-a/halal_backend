package lib

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/takeru-a/halal_backend/models"
	dbmodels "github.com/takeru-a/halal_backend/db/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type ShopSQLRepo struct {
	db *sql.DB
}

type ShopRepository interface {
	Save(ctx *gin.Context, shop models.Shop) error
	FindAll(ctx *gin.Context) (*[]models.Shop, error)
	FindById(ctx *gin.Context, id string) (*models.Shop, error)
	DeleteById(ctx *gin.Context, id string) error
	Update(ctx *gin.Context, shop models.Shop) error
}

func NewShopRepository(db *sql.DB) ShopRepository {
	return &ShopSQLRepo{
		db: db,
	}
}

func (repo *ShopSQLRepo) DeleteById(ctx *gin.Context, id string) error {
	m, err := dbmodels.FindShop(ctx, repo.db, id)
	if err == sql.ErrNoRows {
		return err
	}
	if err != nil {
		return err
	}
	_, delErr := m.Delete(ctx, repo.db)
	if delErr != nil {
		return err
	}
	return nil
}

func (repo *ShopSQLRepo) Update(ctx *gin.Context, shop models.Shop) error {
	updTarget := dbmodels.Shop{
		ID:	shop.ID,
		Name:	shop.Name,
		Introduction:	shop.Introduction,
		Location:	shop.Location,
		Level:	shop.Level,
		Score: shop.Score,
	}

	_, err := updTarget.Update(ctx, repo.db, boil.Infer())
	if err != nil {
		return err
	}
	return nil
}

func (repo *ShopSQLRepo) Save(ctx *gin.Context, shop models.Shop) error {
	saveTarget :=  dbmodels.Shop{
		ID:     shop.ID,
		Name:  shop.Name,
		Introduction: shop.Introduction,
		Location:  shop.Location,
		Level: shop.Level,
		Score: shop.Score,
	}

	err := saveTarget.Insert(ctx, repo.db, boil.Infer())
	if err != nil {
		return err
	}
	return nil
}

func (repo *ShopSQLRepo) FindAll(ctx *gin.Context) (*[]models.Shop, error) {

	m, err := dbmodels.Shops().All(ctx, repo.db)
	if err != nil {
		return nil, err
	}

	result := make([]models.Shop, len(m))
	for i, s := range m {
		shop, _ := models.NewShop(s.ID, s.Name, s.Introduction, s.Location, s.Level, s.Score)
		result[i] = *shop
	}
	return &result, nil
}

func (repo *ShopSQLRepo) FindById(ctx *gin.Context, id string) (*models.Shop, error) {
	m, err := dbmodels.FindShop(ctx, repo.db, id)
	if err == sql.ErrNoRows {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	shop, _ := models.NewShop(m.ID, m.Name, m.Introduction, m.Location, m.Level, m.Score)
	return shop, nil
}
