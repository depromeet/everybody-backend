package main

import (
	"fmt"
	"strings"

	"github.com/ashwanthkumar/slack-go-webhook"
	"github.com/depromeet/everybody-backend/rest-api/config"
	log "github.com/sirupsen/logrus"
)

type loggerHook struct{}

func (h *loggerHook) Levels() []log.Level {
	return []log.Level{log.ErrorLevel}
}

func (h *loggerHook) Fire(entry *log.Entry) error {
	// 비동기적 수행
	go func() {
		defer recover()
		s, err := entry.String()
		if err != nil {
			panic(err)
		}
		if config.Config.ErrorLog.Slack.Enabled {
			title := "눈바디 에러 로그"
			// \\n 텍스트를 \n로 변경해야 슬랙에서 가독성 좋게 출력됨.
			text := strings.ReplaceAll(s, "\\n", "\n")
			color := "#FF0000"
			payload := slack.Payload{
				Username:  config.Config.ErrorLog.Slack.Username,
				Channel:   config.Config.ErrorLog.Slack.Channel,
				IconEmoji: config.Config.ErrorLog.Slack.IconEmoji,
				Attachments: []slack.Attachment{slack.Attachment{
					Title: &title,
					Text:  &text,
					Color: &color,
				}},
			}
			err := slack.Send(config.Config.ErrorLog.Slack.Webhook, "", payload)
			if len(err) > 0 {
				fmt.Printf("error: %s\n", err)
			}
		}
	}()

	return nil
}

func initializeLogger() {
	log.SetReportCaller(true)
	log.AddHook(&loggerHook{})
}
