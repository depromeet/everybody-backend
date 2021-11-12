package controller

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/labstack/echo/v4"
)

type IndexController struct{}

func (c *IndexController) Index(ctx echo.Context) error {
	log.Debug("Index OK...")
	ctx.Response().Header().Set("Content-Type", "text/html")
	ctx.Response().Header().Set("Health-Checked-Time", time.Now().Format(time.RFC3339))
	return ctx.String(200, fmt.Sprint(`
<html>
<head>
<meta charset="utf-8" />
<meta property="og:title" content="눈바디::한눈에 보이는 체형 변화!" />
<meta property="og:type" content="website" />
<meta property="og:description" content="천천히 그리고 꾸준히!" />
<meta property="og:image" content="https://everybody-public-drive.s3.ap-northeast-2.amazonaws.com/logo-wide.png" />
</head>
<body>
<h2>디프만 눈바디 서비스 백엔드 API</h2>
</body>
</html>
`))
}
