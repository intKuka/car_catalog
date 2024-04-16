package filters

import "car_catalog/internal/models"

type Filter struct {
	Page      int `json:"page" validate:"required,gt=0"`
	PageLimit int `json:"pageLimit" validate:"required,gt=0"`
	models.Car
}
