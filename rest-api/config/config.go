package config

import (
	"errors"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	Config = &config{}
	viper.AddConfigPath(os.Getenv("EVERYBODY_REST_CONFIG_PATH"))
	env := strings.ToLower(os.Getenv("EVERYBODY_ENVIRONMENT"))
	if len(env) == 0 {
		log.Warningf("어떤 환경을 이용해 서버를 띄울지 선택해주세요. e.g. EVERYBODY_ENVIRONMENT=local")
	}

	viper.SetEnvPrefix("EVERYBODY")
	viper.AutomaticEnv()

	viper.SetConfigName(env)
	if err := viper.ReadInConfig(); err != nil {
		if errors.As(err, &viper.ConfigFileNotFoundError{}) {
			log.Warningf("설정파일을 하나도 찾지 못했습니다. 올바른 환경을 설정하시고, 그에 대한 설정파일을 생성해주세요.")
		} else {
			log.Fatal(err)
		}
	} else {
		log.Infof("%s 환경 대한 설정파일을 발견했습니다.", env)

		if err := viper.Unmarshal(Config); err != nil {
			log.Fatal(err)
		}
	}
	log.Infof("%+v", Config)
}

var Config *config

type config struct {
	Port int
	DB   struct {
		MySQL struct {
			Host         string
			Port         int
			DatabaseName string
			User         string
			Password     string
		}
	}

	AWS struct {
		Profile string
		Region  string
		Bucket  string
	}
	PublicDriveRootURL string
	ImageRootUrl       string
	ImagePublicKeyID   string
	ImagePrivateKey    string
	Push               struct {
		FCM struct {
			ServiceAccountFile string
		}
	}
	NotifyRoutine struct {
		Enabled  bool
		Interval int
	}
	ErrorLog struct {
		Slack struct {
			Enabled   bool
			Webhook   string
			Channel   string
			Username  string
			IconEmoji string
		}
	}
	Feedback struct {
		Slack struct {
			Webhook   string
			Channel   string
			Username  string
			IconEmoji string
		}
	}
}
