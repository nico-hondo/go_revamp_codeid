package services

import (
	"net/http"

	"codeid.revamptwo/models"
	 "codeid.revamptwo/repositories"
	dbc "codeid.revamptwo/repositories/dbContext"

	"github.com/gin-gonic/gin"
)

type SpecialOfferService struct {
	specialOfferRepository *repositories.SpecialOfferRepository
}

func NewSpecialOfferService(specialOfferRepository *repositories.SpecialOfferRepository) *SpecialOfferService {
	return &SpecialOfferService{
		specialOfferRepository: specialOfferRepository,
	}
}

func (cs SpecialOfferService) GetSpecialOffer(ctx *gin.Context, id int64) (*models.SalesSpecialOffer, *models.ResponseError) {
	return cs.specialOfferRepository.GetSpecialOffer(ctx, id)
}

func validateSpecialOffer(specialOfferParams *dbc.CreateSales_special_offerParams) *models.ResponseError {
	if specialOfferParams.SpofID == 0 {
		return &models.ResponseError{
			Message: "Invalid category id",
			Status:  http.StatusBadRequest,
		}
	}

	if specialOfferParams.SpofDescription == "" {
		return &models.ResponseError{
			Message: "Invalid special offer name",
			Status:  http.StatusBadRequest,
		}
	}

	return nil

}
