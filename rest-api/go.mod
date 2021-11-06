module github.com/depromeet/everybody-backend/rest-api

go 1.16

require (
	entgo.io/ent v0.9.1
	firebase.google.com/go/v4 v4.6.0
	github.com/aws/aws-sdk-go v1.41.1
	github.com/go-sql-driver/mysql v1.5.1-0.20200311113236-681ffa848bae
	github.com/gofiber/fiber/v2 v2.19.0
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/viper v1.7.0
	github.com/stretchr/testify v1.7.0
)

replace github.com/pkg/errors v0.9.1 => github.com/depromeet/errors v0.9.2
