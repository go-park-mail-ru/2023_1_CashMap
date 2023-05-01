package utils

import "github.com/jmoiron/sqlx"

func CloseRows(rows *sqlx.Rows) {
	if rows != nil {
		_ = rows.Close()
	}
}
