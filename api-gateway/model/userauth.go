package model

import (
	"errors"

	log "github.com/sirupsen/logrus"

	"github.com/depromeet/everybody-backend/api-gateway/util"
)

type UserAuth struct {
	UserId   int    `json:"user_id"`
	SocialId string `json:"social_id"`
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

	return &UserAuth{
		UserId:   userId,
		Password: password,
	}, nil
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

func GetUserAuthBySocialId(sid string) (*UserAuth, error) {
	sqlStatement := "SELECT user_id, social_id, password FROM UserAuth WHERE social_id = ?"
	conn := util.CreateDBConn()
	defer conn.Close()

	var userId int
	var password string
	var socialId string
	err := conn.QueryRow(sqlStatement, sid).Scan(&userId, &socialId, &password)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &UserAuth{userId, socialId, password}, nil
}

func SetUserAuthWithSocialId(userId int, socialId string) error {
	sqlStatement := "UPDATE UserAuth SET social_id = ? WHERE user_id = ?"
	conn := util.CreateDBConn()
	defer conn.Close()

	result, err := conn.Exec(sqlStatement, socialId, userId)
	if err != nil {
		log.Error("SetUserAuthWithSocialId -> ", err)
		return err
	}
	n, err := result.RowsAffected()
	if err != nil {
		log.Error("SetUserAuthWithSocialId -> ", err)
		return err
	}

	if n < 1 {
		log.Info("UserAuth 테이블에 해당하는 row가 없습니다.")
		return errors.New("해당하는 유저 정보가 없습니다")
	}

	return nil
}
