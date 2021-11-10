module github.com/depromeet/everybody-backend/rest-api

go 1.16

require (
	entgo.io/ent v0.9.1
	firebase.google.com/go/v4 v4.6.0
	github.com/ashwanthkumar/slack-go-webhook v0.0.0-20200209025033-430dd4e66960
	github.com/aws/aws-sdk-go v1.41.1
	github.com/elazarl/goproxy v0.0.0-20210801061803-8e322dfb79c4 // indirect
	github.com/go-sql-driver/mysql v1.5.1-0.20200311113236-681ffa848bae
	github.com/gofiber/fiber/v2 v2.19.0
	github.com/parnurzeal/gorequest v0.2.16 // indirect
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/viper v1.7.0
	github.com/stretchr/testify v1.7.0
	moul.io/http2curl v1.0.0 // indirect
)

replace github.com/pkg/errors v0.9.1 => github.com/depromeet/errors v0.9.2
