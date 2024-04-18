package helpers

import (
	"car_catalog/internal/models"
	"car_catalog/internal/models/filters"
	"fmt"
	"strconv"
	"strings"
)

// TODO: consider to add JOIN
func FindByFilterSql(f *filters.Filter) string {

	sql := `SELECT * FROM Cars`
	var filterFields []string // stores which filters must be applied

	if f.RegNum != "" {
		filterFields = append(filterFields, "regnum = '"+f.RegNum+"'")
	}

	if f.Mark != "" {
		filterFields = append(filterFields, "mark = '"+f.Mark+"'")
	}

	if f.Model != "" {
		filterFields = append(filterFields, "model = '"+f.Model+"'")
	}

	if f.Year != 0 {
		filterFields = append(filterFields, "year = '"+strconv.Itoa(f.Year)+"'")
	}

	if f.Owner.Name != "" {
		filterFields = append(filterFields, "name = '"+f.Owner.Name+"'")
	}

	if f.Owner.Surname != "" {
		filterFields = append(filterFields, "surname = '"+f.Owner.Surname+"'")
	}

	if f.Owner.Patronymic != "" {
		filterFields = append(filterFields, "patronymic = '"+f.Owner.Patronymic+"'")
	}

	// TODO: Grab pageLimit from config
	paging := fmt.Sprintf("LIMIT %d OFFSET %d", f.PageLimit, (f.Page-1)*f.PageLimit)

	if len(filterFields) > 0 {
		sqlWhere := strings.Join(filterFields, " AND ")
		sql = strings.Join([]string{sql, "WHERE", sqlWhere, paging}, " ")
	} else {
		sql = strings.Join([]string{sql, paging}, " ")
	}

	return sql
}

func UpdateFieldsByIdSql(c *models.Car) string {
	sql := `UPDATE Cars
	SET %s
	WHERE id = $1;`

	var setFields []string // stores which fields must be updated

	if c.RegNum != "" {
		setFields = append(setFields, "regnum = '"+c.RegNum+"'")
	}

	if c.Mark != "" {
		setFields = append(setFields, "mark = '"+c.Mark+"'")
	}

	if c.Model != "" {
		setFields = append(setFields, "model = '"+c.Model+"'")
	}

	if c.Year != 0 {
		setFields = append(setFields, "year = '"+strconv.Itoa(c.Year)+"'")
	}

	if c.Owner.Name != "" {
		setFields = append(setFields, "name = '"+c.Owner.Name+"'")
	}

	if c.Owner.Surname != "" {
		setFields = append(setFields, "surname = '"+c.Owner.Surname+"'")
	}

	if c.Owner.Patronymic != "" {
		setFields = append(setFields, "patronymic = '"+c.Owner.Patronymic+"'")
	}

	setString := strings.Join(setFields, ", ")

	return fmt.Sprintf(sql, setString)
}
