package model

import (
	log "github.com/sirupsen/logrus"

	"github.com/depromeet/everybody-backend/api-gateway/util"
)

type UserAuth struct {
	UserId   uint64 `json:"user_id"`
	Password string `json:"password"`
}

func GetUserAuth(u uint64) UserAuth {
	sqlStatement := "SELECT user_id, password FROM UserAuth WHERE user_id = ?"
	conn := util.CreateDBConn()
	defer conn.Close()

	var userId uint64
	var password string
	err := conn.QueryRow(sqlStatement, u).Scan(&userId, &password)
	if err != nil {
		log.Error(err)
	}

	return UserAuth{userId, password}
}

func SetUserAuth(ua UserAuth) error {
	sqlStatement := "INSERT INTO UserAuth(user_id, password) VALUES(?, ?)"
	conn := util.CreateDBConn()
	defer conn.Close()

	result, err := conn.Exec(sqlStatement, ua.UserId, ua.Password)
	if err != nil {
		log.Fatal(err)
		return err
	}
	n, err := result.RowsAffected()
	if n != 1 {
		return err
	}

	return nil
}
