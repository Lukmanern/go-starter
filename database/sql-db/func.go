package database

import "errors"

func TryPingDB() error {
	return nil
}

func IsTableExist(tableName string) bool {
	return true
}

func AddTable(tableName string) error {
	if IsTableExist(tableName) {
		return errors.New("table name is exist")
	}

	return nil
}
