package initialize

import (
	"Go-000/Week02/config"
	"database/sql"
	"fmt"
)

func InitDB() (*sql.DB, error)  {
	m := config.GetDBConfiguration()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True",
		m.UserName, m.Password, m.Host, m.Port, m.DBName)
	db, err := sql.Open("mysql", dsn)
	return db, err
}
