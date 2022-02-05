package model

import (
	"errors"

	"github.com/depromeet/everybody-backend/api-gateway/util"
	log "github.com/sirupsen/logrus"
)

type UserAuth struct {
	UserId     int
	SocialId   string
	SocialKind string
	Password   string
	Status     string
}

var userAuthStatus = []string{"active", "inactive"}

func GetUserAuth(u int) (*UserAuth, error) {
	sqlStatement := "SELECT user_id, password, status FROM UserAuth WHERE user_id = ?"
	conn := util.CreateDBConn()
	defer conn.Close()

	userAuth := UserAuth{}
	err := conn.QueryRow(sqlStatement, u).Scan(&userAuth.UserId, &userAuth.Password, &userAuth.Status)
	if err != nil {
		log.Error("GetUserAuth -> ", err)
		return nil, err
	}

	if userAuth.Status != userAuthStatus[0] {
		log.Error("회원탈퇴한 유저입니다")
		return nil, errors.New("존재하지 않는 유저입니다")
	}
	return &userAuth, nil
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

func GetUserAuthBySocialId(sid, sKind string) (*UserAuth, error) {
	sqlStatement := "SELECT user_id, social_id, social_kind, password, status FROM UserAuth WHERE social_id = ? and social_kind = ?"
	conn := util.CreateDBConn()
	defer conn.Close()

	userAuth := UserAuth{}
	err := conn.QueryRow(sqlStatement, sid, sKind).Scan(&userAuth.UserId, &userAuth.SocialId, &userAuth.SocialKind, &userAuth.Password, &userAuth.Status)
	if err != nil {
		log.Error("GetUserAuthBySocialid -> ", err)
		return nil, err
	}

	if userAuth.Status == userAuthStatus[1] {
		log.Error("회원탈퇴한 유저입니다")
		return nil, errors.New("존재하지 않는 유저입니다")
	}

	return &userAuth, nil
}

func SetUserAuthWithSocial(userId int, sid, sKind string) error {
	sqlStatement := "UPDATE UserAuth SET social_id = ?, social_kind = ? WHERE user_id = ?"
	conn := util.CreateDBConn()
	defer conn.Close()

	result, err := conn.Exec(sqlStatement, sid, sKind, userId)
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
		log.Error("SetUserAuthWithSocialId 실패-> ", userId)
		return errors.New("해당하는 유저 정보가 없습니다")
	}

	return nil
}

func SetUserAuthStatus(userId int) error {
	sqlStatement := "UPDATE UserAuth SET status = ? WHERE user_id = ?"
	conn := util.CreateDBConn()
	defer conn.Close()

	result, err := conn.Exec(sqlStatement, userAuthStatus[1], userId)
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
		log.Error("SetUserAuthWithSocialId 실패-> ", userId)
		return errors.New("해당하는 유저 정보가 없습니다")
	}

	return nil
}
