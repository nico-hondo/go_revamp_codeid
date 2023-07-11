package services

import (
	"net/http"

	"codeid.revamptwo/models"
	"codeid.revamptwo/repositories"
	"codeid.revamptwo/repositories/dbContext"
	"github.com/gin-gonic/gin"
)

type DepartmentService struct {
	departmentRepository *repositories.DepartmentRepository
}

func NewDepartmentService(departmentRepository *repositories.DepartmentRepository) *DepartmentService {
	return &DepartmentService{
		// struct				parameter
		departmentRepository: departmentRepository,
	}
}

// method
func (cs DepartmentService) GetListDepartment(ctx *gin.Context) ([]*models.HrDepartment, *models.ResponseError) {
	return cs.departmentRepository.GetListDepartment(ctx)
}

func (cs DepartmentService) GetDepartment(ctx *gin.Context, id int64) (*models.HrDepartment, *models.ResponseError) {
	return cs.departmentRepository.GetDepartment(ctx, id)
}

// CREATE
func (cs DepartmentService) CreateDepartment(ctx *gin.Context, departmentParams *dbContext.CreateDepartmentParams) (*models.HrDepartment, *models.ResponseError) {
	responseErr := validateDepartment(departmentParams)
	if responseErr != nil {
		return nil, responseErr
	}

	return cs.departmentRepository.CreateDepartment(ctx, departmentParams)
}

// UPDATE
func (cs DepartmentService) UpdateDepartment(ctx *gin.Context, departmentParams *dbContext.CreateDepartmentParams, id int64) *models.ResponseError {
	responseErr := validateDepartment(departmentParams)
	if responseErr != nil {
		return responseErr
	}

	return cs.departmentRepository.UpdateDepartment(ctx, departmentParams)
}

// DELETE
func (cs DepartmentService) DeleteDepartment(ctx *gin.Context, id int64) *models.ResponseError {
	return cs.departmentRepository.DeleteDepartment(ctx, id)
}

// method validation, untuk yang diatas ini
func validateDepartment(departmentParams *dbContext.CreateDepartmentParams) *models.ResponseError {
	if departmentParams.DeptID == 0 {
		return &models.ResponseError{
			Message: "Invalid category id",
			Status:  http.StatusBadRequest,
		}
	}

	if departmentParams.DeptName == "" {
		return &models.ResponseError{
			Message: "Invalid category name",
			Status:  http.StatusBadRequest,
		}
	}

	return nil
}
