package dao

import (
	"Go-000/Week02/config"
	"Go-000/Week02/models"
	"database/sql"
	"github.com/pkg/errors"
)

var (
	NotFoundErr = errors.New("record not found")
)

func GetUserById(userId int) (*models.User, error)  {
	var user models.User
	db := config.DataBase
	err := db.QueryRow("select id, name, age from user where id=?", userId).Scan(&user.ID, &user.Name, &user.Age)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// 直接抛给上层处理
			return nil, errors.Wrap(NotFoundErr, "sql failed query")
		}
		return nil, errors.Wrap(err, "sql failed query")
	}
	return &user, nil
}
