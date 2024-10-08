package data

import (
	"database/sql"
	"sync"
)

var (
	data *Data
	once sync.Once
)

type Data struct {
	DB *sql.DB
}

func initDB() {
	db, err := getConnection()
	if err != nil {
		panic(err)
	}

	err = MakeMigration(db)
	if err != nil {
		panic(err)
	}

	data = &Data{
		DB: db,
	}
}

func New() *Data {
	once.Do(initDB)

	return data
}

func Close() error {
	if data == nil {
		return nil
	}

	return data.DB.Close()
}
