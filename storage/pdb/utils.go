package pdb

import (
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"idon.com/global"
	"strings"
)

const (
	VALUE_BETTER = 'b'
	VALUE_LESS   = 'l'
	VALUE_EQUALS = 'e'
	VALUE_LIKE   = 'k'
)

type FieldsMap = map[string]string

func getLimitSqlQuery(ol *global.OptionsList) string {
	if ol.Limit > 0 {
		return fmt.Sprintf("LIMIT %d", ol.Limit)
	}

	return ""
}

func getOffsetSqlQuery(ol *global.OptionsList) string {
	if ol.Offset > 0 {
		return fmt.Sprintf("OFFSET %d", ol.Limit)
	}

	return ""
}

func getSortsSqlQuery(ol *global.OptionsList, fm FieldsMap) string {
	if (len(ol.Sorts) == 1 && ol.Sorts[0] == "") || len(ol.Sorts) == 0 {
		return ""
	}

	sortsSqlQuery := strings.Builder{}
	sortsSqlQuery.WriteString("ORDER BY")

	var err error
	for _, field := range ol.Sorts {
		value, ok := fm[field[:len(field)-1]]
		if !ok {
			continue
		}

		if strings.HasSuffix(field, "-") {
			_, err = fmt.Fprintf(&sortsSqlQuery, " %s DESC,", value)
		} else {
			_, err = fmt.Fprintf(&sortsSqlQuery, " %s ASC,", value)
		}

		if err != nil {
			break
		}
	}

	if err != nil {
		log.Errorf("GetSortsSqlQuery err: %v", err)
		return ""
	}

	resStr := sortsSqlQuery.String()
	if resStr == "ORDER BY" {
		return ""
	}

	fmt.Println("Error!!! - ", resStr)

	return resStr[:len(resStr)-1]
}

func getFiltersSqlQuery(ol *global.OptionsList, fm FieldsMap, isStart ...bool) string {
	if len(ol.Filters) == 0 {
		return ""
	}

	filtersSqlQuery := strings.Builder{}

	if len(isStart) > 0 && isStart[0] {
		filtersSqlQuery.WriteString("WHERE")
	} else {
		filtersSqlQuery.WriteString("AND")
	}

	var err error
	for k, v := range ol.Filters {
		if len(k) < 2 {
			continue
		}

		kField, ok := fm[k[1:]]
		if !ok {
			continue
		}

		switch k[0] {
		case VALUE_EQUALS:
			_, err = fmt.Fprintf(&filtersSqlQuery, " %s = '%s' AND", kField, v)
		case VALUE_BETTER:
			_, err = fmt.Fprintf(&filtersSqlQuery, " %s >= %s AND", kField, v)
		case VALUE_LESS:
			_, err = fmt.Fprintf(&filtersSqlQuery, " %s <= %s AND", kField, v)
		case VALUE_LIKE:
			_, err = fmt.Fprintf(&filtersSqlQuery, " %s LIKE '%%%s%%' AND", kField, v)
		}

		if err != nil {
			break
		}
	}

	if err != nil {
		log.Errorf("GetFiltersSqlQuery err: %v", err)
		return ""
	}

	resStr := filtersSqlQuery.String()

	if resStr == "WHERE" || resStr == "AND" {
		return ""
	}

	return resStr[:len(resStr)-4]
}

func getQueryWithOptions(querySql string, ol *global.OptionsList, filedMap FieldsMap) string {
	// TODO: подвержено SQL-иньъекциям, придумать способ иначе констурировать запрос
	isStart := true
	if strings.Contains(querySql, "WHERE ") {
		isStart = false
	}

	// добавляем фильтры
	filtersSql := getFiltersSqlQuery(ol, filedMap, isStart)

	// добавляем сортировки
	sortsSql := getSortsSqlQuery(ol, filedMap)

	// добавляем ограничение
	limitsSql := getLimitSqlQuery(ol)

	// добавляем смещение
	offsetSql := getOffsetSqlQuery(ol)

	resQuery := fmt.Sprintf(
		"%s\n%s\n%s\n%s\n%s;",
		querySql,
		filtersSql,
		sortsSql,
		limitsSql,
		offsetSql,
	)

	return resQuery
}
