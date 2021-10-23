package util

import (
	"database/sql"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"

	"github.com/depromeet/everybody-backend/api-gateway/config"
)

func CreateDBConn() *sql.DB {
	db, err := sql.Open("mysql", config.Config.DB.MySQL.User+":"+config.Config.DB.MySQL.Password+"@tcp("+config.Config.DB.MySQL.Host+":"+strconv.Itoa(config.Config.DB.MySQL.Port)+")/"+config.Config.DB.MySQL.DatabaseName)
	if err != nil {
		log.Fatal(err.Error())
	}
	//defer db.Close()

	err = db.Ping() // make sure connection is available
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Info("db is connected")
	return db
}
