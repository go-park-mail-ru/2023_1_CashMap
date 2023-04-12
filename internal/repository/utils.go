package repository

import (
	"errors"
	"fmt"
	"github.com/fatih/structs"
	"github.com/jmoiron/sqlx"
	"reflect"
	"strings"
)

func UpdateTable(db *sqlx.DB, tableName string, conditionField string, conditionValue interface{},
	structToDB map[string]string, entity interface{}) error {
	query := "UPDATE " + tableName + " SET "
	values := make([]interface{}, 0, len(structToDB))
	for name, value := range structs.Map(entity) {
		if !IsNil(value) {
			if field, ok := structToDB[name]; ok {
				query += fmt.Sprintf("%s = $%d, ", field, len(values)+1)
				values = append(values, value)
			}
		}
	}
	if len(values) == 0 {
		return errors.New("rows were not affected")
	}

	values = append(values, conditionValue)

	query = strings.TrimRight(query, ", ")

	query += fmt.Sprintf(" WHERE %s = $%d;", conditionField, len(values))
	_, err := db.Exec(query, values...)
	if err != nil {
		return err
	}

	return nil
}

func IsNil(value interface{}) bool {
	if value == nil {
		return true
	}
	val := reflect.ValueOf(value)
	switch val.Kind() {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Pointer, reflect.UnsafePointer, reflect.Interface, reflect.Slice:
		return val.IsNil()
	}
	return false
}
