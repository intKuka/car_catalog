package queries

import (
	"car_catalog/internal/models/filters"
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// TODO: consider to add JOIN
func WriteSqlWithFilter(c *gin.Context, f *filters.Filter) string {

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
		filterFields = append(filterFields, "year = '"+strconv.Itoa(f.Year))
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

func WriteSqlWithUpdate() string {
	sql := `
		UPDATE Products
		SET Price = Price + 3000;`

	return sql
}
