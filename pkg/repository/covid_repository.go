package repository

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/nutthanonn/covid-data/pkg/domain/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type covidRepository struct {
	db *mongo.Database
}

type CovidRepository interface {
	CallData() (*models.CovidDataResponse, error)
	GetData() ([]*models.CovidModel, error)
	GetCaseByYear() ([]*models.CovidByYear, error)
	GetCaseByJob() ([]*models.CovidByJob, error)
}

func NewCovidRepository(db *mongo.Database) CovidRepository {
	return &covidRepository{db}
}

func (c *covidRepository) CallData() (*models.CovidDataResponse, error) {
	/*
		https://covid19.ddc.moph.go.th/api/Cases/round-1to2-line-lists
		https://covid19.ddc.moph.go.th/api/Cases/data-round-3-y21-line-lists-by-province
		https://covid19.ddc.moph.go.th/api/Cases/round-4-line-lists
	*/

	resp, err := http.Get("https://covid19.ddc.moph.go.th/api/Cases/round-4-line-lists")

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var covidResponse models.CovidDataResponse
	if err := json.NewDecoder(resp.Body).Decode(&covidResponse); err != nil {
		return nil, err
	}

	collection := c.db.Collection("data")
	var dataToInsert []interface{}

	for _, d := range covidResponse.Data {
		dataToInsert = append(dataToInsert, d)
	}

	if _, err := collection.InsertMany(context.TODO(), dataToInsert); err != nil {
		return nil, err
	}

	return &covidResponse, nil
}

func (c *covidRepository) GetData() ([]*models.CovidModel, error) {
	var covidResponse []*models.CovidModel
	collection := c.db.Collection("data")

	cur, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		var elem models.CovidModel
		if err := cur.Decode(&elem); err != nil {
			return nil, err
		}
		covidResponse = append(covidResponse, &elem)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return covidResponse, nil
}

func (c *covidRepository) GetCaseByYear() ([]*models.CovidByYear, error) {
	covidResponse, err := c.GetData()
	if err != nil {
		return nil, err
	}

	// enter code under this

	var covidCaseByYear []*models.CovidByYear
	year_case := make(map[int]int)

	for _, covid_case := range covidResponse {
		if _, ok := year_case[covid_case.Year]; ok {
			year_case[covid_case.Year] = year_case[covid_case.Year] + 1
		} else {
			year_case[covid_case.Year] = 1
		}
	}

	for key, value := range year_case {
		covidCaseByYear = append(covidCaseByYear, &models.CovidByYear{Year: key, Count: value})
	}

	return covidCaseByYear, nil
}

func (c *covidRepository) GetCaseByJob() ([]*models.CovidByJob, error) {
	covidResponse, err := c.GetData()
	if err != nil {
		return nil, err
	}

	var covidCaseByJob []*models.CovidByJob
	job := make(map[string]int)

	for _, covid_case := range covidResponse {
		if _, ok := job[covid_case.Job]; ok {
			job[covid_case.Job] = job[covid_case.Job] + 1
		} else {
			job[covid_case.Job] = 1
		}
	}

	for key, value := range job {
		if key == "" {
			covidCaseByJob = append(covidCaseByJob, &models.CovidByJob{Job: "N/A", Count: value})
		} else {
			covidCaseByJob = append(covidCaseByJob, &models.CovidByJob{Job: key, Count: value})
		}
	}

	return covidCaseByJob, nil
}
