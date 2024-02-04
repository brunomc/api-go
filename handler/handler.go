package handler

import (
	"fmt"
	"net/http"

	"github.com/brunomc/api-go/config"
	"github.com/brunomc/api-go/schemas"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitializeHandler() {
	logger = config.GetLogger("handler")
	db = config.GetSQLite()
}
func CreateOpeningHandler(ctx *gin.Context) {
	request := CreateOpeningRequest{}
	ctx.BindJSON(&request)
	if err := request.Validate(); err != nil {
		logger.Errorf("validation error %v", err.Error())
		sendError(ctx, 400, err.Error())
		return
	}
	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Link:     request.Link,
		Remote:   *request.Remote,
		Salary:   request.Salary,
	}
	if err := db.Create(&opening).Error; err != nil {
		sendError(ctx, 400, err.Error())
		return
	}
	sendSuccess(ctx, "create-opening", opening)
}
func ShowOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	opening := schemas.Opening{}
	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusBadRequest, fmt.Errorf("opening with id %s not found", id).Error())
		return
	}
	sendSuccess(ctx, "show-opening", opening)

}
func DeleteOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	opening := schemas.Opening{}
	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusBadRequest, fmt.Errorf("opening with id %s not found", id).Error())
		return
	}
	if err := db.Delete(&opening).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Errorf("error on delete opening with id %s", id).Error())
		return
	}
	sendSuccess(ctx, "delete-opening", opening)
}
func ListOpeningHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}
	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "erro on list openings")
		return
	}
	sendSuccess(ctx, "list-openings", openings)
}
func UpdateOpeningHandler(ctx *gin.Context) {
	request := UpdateOpeningRequest{}
	ctx.BindJSON(&request)
	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	opening := schemas.Opening{}
	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusBadRequest, fmt.Errorf("opening with id %s not found", id).Error())
		return
	}
	if request.Role != "" {
		opening.Role = request.Role
	}
	if request.Company != "" {
		opening.Company = request.Company
	}
	if request.Location != "" {
		opening.Location = request.Location
	}
	if request.Link != "" {
		opening.Link = request.Link
	}
	if request.Remote != nil {
		opening.Remote = *request.Remote
	}
	if request.Salary > 0 {
		opening.Salary = request.Salary
	}
	if err := db.Save(&opening).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Errorf("error on update opening with id: %s", id).Error())
		return
	}
	sendSuccess(ctx, "update-opening", opening)

}
