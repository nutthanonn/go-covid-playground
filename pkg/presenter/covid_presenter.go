package presenter

import (
	"github.com/gin-gonic/gin"
	"github.com/nutthanonn/covid-data/pkg/domain/models"
)

type covidPresenter struct {
}

type CovidPresenter interface {
	CovidSuccessResponse(data []*models.CovidModel) gin.H
	CovidErrorResponse(err error) gin.H
}

func NewCovidPresenter() CovidPresenter {
	return &covidPresenter{}
}

func (covidPresenter *covidPresenter) CovidSuccessResponse(data []*models.CovidModel) gin.H {
	return gin.H{
		"error": nil,
		"data":  data,
	}
}

func (covidPresenter *covidPresenter) CovidErrorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
		"data":  nil,
	}
}
