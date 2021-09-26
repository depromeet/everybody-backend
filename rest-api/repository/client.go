package repository

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"fmt"
	"github.com/depromeet/everybody-backend/rest-api/config"
	"github.com/depromeet/everybody-backend/rest-api/ent"
	"github.com/depromeet/everybody-backend/rest-api/ent/migrate"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"time"
)

func Connect() *ent.Client {
	drv, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true",
		config.Config.DB.MySQL.User,
		config.Config.DB.MySQL.Password,
		config.Config.DB.MySQL.Host,
		config.Config.DB.MySQL.DatabaseName,
	))
	if err != nil {
		log.Fatal(err)
	}

	db := drv.DB()
	// TODO: 몇으로 할까요
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(50)
	db.SetConnMaxLifetime(time.Minute * 30)

	// ping
	conn, err := db.Conn(context.TODO())
	if err != nil {
		log.Panic(err)
	}
	defer conn.Close()

	ent.Debug()
	ent.Log(func(i ...interface{}) {
		log.Warning(i...)
	})

	client := ent.NewClient(ent.Driver(drv))
	err = client.Schema.Create(
		context.TODO(),
		migrate.WithDropIndex(true),
		migrate.WithForeignKeys(true),
		//migrate.WithDropColumn(true), // 데이터 날아갈 수도 있음...
	)

	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	log.Info("데이터베이스 연결 완료")

	return client
}
