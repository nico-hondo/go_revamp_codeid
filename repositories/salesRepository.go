package repositories

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SpecialOfferRepository struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewSpecialOfferRepository(dbHandler *sql.DB) *SpecialOfferRepository {
	return &SpecialOfferRepository{
		dbHandler: dbHandler,
	}
}

func (cr SpecialOfferRepository) GetSpecialOffer(ctx *gin.Context, id int64) (*models.SpecialOffer, *models.ResponseError) {

	store := dbContext.New(cr.dbHandler)
	specialoffer, err := store.GetSpecialOffer(ctx, int16(id))

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &specialoffer, nil
}
