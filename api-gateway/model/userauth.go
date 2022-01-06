package model

import (
	log "github.com/sirupsen/logrus"

	"github.com/depromeet/everybody-backend/api-gateway/util"
)

type UserAuth struct {
	UserId   int    `json:"user_id"`
	Password string `json:"password"`
}

func GetUserAuth(u int) (*UserAuth, error) {
	sqlStatement := "SELECT user_id, password FROM UserAuth WHERE user_id = ?"
	conn := util.CreateDBConn()
	defer conn.Close()

	var userId int
	var password string
	err := conn.QueryRow(sqlStatement, u).Scan(&userId, &password)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &UserAuth{userId, password}, nil
}

func SetUserAuth(ua UserAuth) error {
	sqlStatement := "INSERT INTO UserAuth(user_id, password) VALUES(?, ?)"
	conn := util.CreateDBConn()
	defer conn.Close()

	result, err := conn.Exec(sqlStatement, ua.UserId, ua.Password)
	if err != nil {
		log.Fatal("SetUserAuth -> ", err)
	}
	n, err := result.RowsAffected()
	if n != int64(1) || err != nil {
		log.Fatal("SetUserAuth -> ", err)
	}

	return nil
}
