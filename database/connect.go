package database

import (
	"database/sql" //Thao tác với SQL
	"log"
	"time" //Xử lý thời gian

	_ "github.com/go-sql-driver/mysql" //Tạo driver kết nối mysql
)

func DbConnection() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/go_admin")
	if err != nil {
		log.Printf("Error %s when opening DBn", err)
		return nil, err
	}

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(20)
	db.SetConnMaxLifetime(time.Minute * 5)
	return db, nil
}
