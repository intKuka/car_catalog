package controller

import (
	"car_catalog/cmd/initializers"
	"car_catalog/internal/helpers/validators"
	"car_catalog/internal/models"
	"car_catalog/internal/models/filters"
	"car_catalog/internal/storage/postgres"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5"
)

// TODO: write global error handler, middleware mb
// TODO: move validator to somewhere

// ListCars godoc
// @Summary      List cars
// @Description  get cars
// @Tags         cars
// @Accept       json
// @Produce      json
// @Success      200  {array}   models.Car
// @Router       /cars [get]
func (c *Controller) ListCars(ctx *gin.Context) {
	var filter filters.Filter

	if err := ctx.BindJSON(&filter); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid filter values"})
		return
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(&filter); err != nil {
		errs := err.(validator.ValidationErrors)

		// TODO: compact into single json
		for _, fieldErr := range errs {
			fmt.Println(validators.MsgForTag(fieldErr))
		}
		return
	}

	cars, err := postgres.GetAllCars(ctx, &filter)
	if err != nil {
		initializers.Log.Error(err.Error())
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, &cars)
}

func (c *Controller) UpdateCar(ctx *gin.Context) {
	var car models.Car
	id := ctx.Param("id")

	if err := ctx.BindJSON(&car); err != nil {
		initializers.Log.Error(err.Error())
		ctx.Status(http.StatusBadRequest)
		return
	}

	if err := postgres.UpdateCarById(ctx, id, &car); err != nil {
		initializers.Log.Error(err.Error())
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			ctx.Status(http.StatusNotFound)
		default:
			ctx.Status(http.StatusInternalServerError)
		}
		return
	}

	ctx.Status(http.StatusOK)
}

func (c *Controller) DeleteCar(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := postgres.DeleteCarById(ctx, id); err != nil {
		initializers.Log.Error(err.Error())
		ctx.Status(http.StatusNotFound)
	}
}
