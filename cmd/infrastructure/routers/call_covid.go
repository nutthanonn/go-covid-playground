package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nutthanonn/covid-data/pkg/presenter"
	"github.com/nutthanonn/covid-data/pkg/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

func CallCovid(api *gin.RouterGroup, quickStart *mongo.Database) {
	api.GET("/v1/covid", func(c *gin.Context) {
		repository := repository.NewCovidRepository(quickStart)

		covid, err := repository.CallData()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
				"data":  nil,
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"error": nil,
			"data":  covid.Data,
		})
	})

	api.GET("/v1/covid/get", func(c *gin.Context) {
		repository := repository.NewCovidRepository(quickStart)
		p := presenter.NewCovidPresenter()

		covid, err := repository.GetData()

		if err != nil {
			c.JSON(http.StatusInternalServerError, p.CovidErrorResponse(err))
		}

		c.JSON(http.StatusOK, p.CovidSuccessResponse(covid))
	})

	api.GET("/v1/covid/case", func(c *gin.Context) {
		repository := repository.NewCovidRepository(quickStart)
		p := presenter.NewCovidPresenter()

		covid, err := repository.GetCaseByYear()

		if err != nil {
			c.JSON(http.StatusInternalServerError, p.CovidErrorResponse(err))
		}

		c.JSON(http.StatusOK, gin.H{
			"data":  covid,
			"erorr": nil,
		})
	})

	api.GET("/v1/covid/job", func(c *gin.Context) {
		repository := repository.NewCovidRepository(quickStart)
		p := presenter.NewCovidPresenter()

		covid, err := repository.GetCaseByJob()

		if err != nil {
			c.JSON(http.StatusInternalServerError, p.CovidErrorResponse(err))
		}

		c.JSON(http.StatusOK, gin.H{
			"data":  covid,
			"erorr": nil,
		})
	})
}
