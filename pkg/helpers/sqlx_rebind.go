package helpers

import "github.com/jmoiron/sqlx"

func RebindPostgre(query string, args ...interface{}) (*string, []interface{}, error) {
	queryIn, arguments, err := sqlx.In(query, args...)
	if err != nil {
		return nil, nil, err
	}
	queryBind := sqlx.Rebind(sqlx.DOLLAR, queryIn)

	return &queryBind, arguments, nil
}
