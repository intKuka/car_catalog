package postgres

import (
	"car_catalog/cmd/initializers"
	"car_catalog/internal/models"
	"car_catalog/internal/models/filters"
	"car_catalog/internal/storage/postgres/helpers"
	"math/rand"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

const sqlLog = "SQL: "

func GetAllCars(c *gin.Context, f *filters.Filter) (*[]models.Car, error) {
	sql := helpers.FindByFilterSql(f)
	initializers.Log.Debug(sqlLog + sql)

	rows, err := initializers.DB.Query(c, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cars []models.Car

	for rows.Next() {
		var car models.Car

		if err := rows.Scan(
			&car.Id,
			&car.RegNum,
			&car.Mark,
			&car.Model,
			&car.Year,
			&car.Owner.Name,
			&car.Owner.Surname,
			&car.Owner.Patronymic,
		); err != nil {
			return nil, err
		}

		cars = append(cars, car)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &cars, nil
}

func CreateCar() {

}

func CreateCarStub(c *gin.Context) {
	const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	modeles := []string{"Q7", "TT3", "It4", "ROgue", "Pollette", "IoT", "Razor"}
	marks := []string{"lada", "nevada", "bmw", "mustang", "porshe", "Tesla"}
	names := []string{"Joh", "Lay", "Troy", "Yappi", "Ulliy", "Tears", "Lon"}

	var cars []models.Car

	sql := `
		INSERT INTO Cars (regnum, mark, model, year, name, surname, patronymic)
		VALUES ($1, $2, $3, $4, $5, $6, $7);`

	batch := &pgx.Batch{}
	iterator, _ := strconv.Atoi(c.Param("num"))

	for i := 0; i < iterator; i++ {
		regnum := make([]byte, 9)
		for i := range regnum {
			regnum[i] = letters[rand.Intn(len(letters))]
		}

		car := models.Car{
			RegNum: string(regnum),
			Mark:   marks[rand.Intn(len(marks))],
			Model:  modeles[rand.Intn(len(modeles))],
			Year:   rand.Intn(150) + 1900,
			Owner: models.People{
				Name:       names[rand.Intn(len(names))],
				Surname:    names[rand.Intn(len(names))],
				Patronymic: names[rand.Intn(len(names))],
			},
		}
		cars = append(cars, car)

		batch.Queue(sql,
			car.RegNum,
			car.Mark,
			car.Model,
			car.Year,
			car.Owner.Name,
			car.Owner.Surname,
			car.Owner.Patronymic)
	}

	results := initializers.DB.SendBatch(c, batch)
	if err := results.Close(); err != nil {
		initializers.Log.Error(err.Error())
		return
	}

	c.JSON(201, gin.H{
		"cars": cars,
	})
}

// TODO: mb write separate struct for update
func UpdateCarById(c *gin.Context, id string, car *models.Car) error {
	sql := "SELECT * FROM Cars WHERE id = " + id

	row := initializers.DB.QueryRow(c, sql)
	if err := row.Scan(); err != nil {
		return err
	}

	sql = helpers.UpdateFieldsByIdSql(car)

	initializers.Log.Debug(sqlLog + sql)

	if _, err := initializers.DB.Exec(c, sql, id); err != nil {
		return err
	}

	return nil
}

func DeleteCarById(c *gin.Context, id string) error {
	sql := "SELECT * FROM Cars WHERE id = " + id

	row := initializers.DB.QueryRow(c, sql)
	if err := row.Scan(); err != nil {
		return err
	}

	sql = "DELETE FROM Cars WHERE id=$1"

	initializers.Log.Debug(sqlLog + sql)

	if _, err := initializers.DB.Exec(c, sql, id); err != nil {
		return err
	}

	return nil
}
