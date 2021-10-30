package config

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"strings"
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
			log.Fatal("%#v", err)
		}
	} else {
		log.Infof("%s 환경 대한 설정파일을 발견했습니다.", env)

		if err := viper.Unmarshal(Config); err != nil {
			log.Fatal(err)
		}
	}
	log.Info(Config)
}

var Config *config

type config struct {
	Port int `yaml:"port"`
	DB   struct {
		MySQL struct {
			Host         string `yaml:"host"`
			Port         int    `yaml:"port"`
			DatabaseName string `yaml:"databaseName"`
			User         string `yaml:"user"`
			Password     string `yaml:"password"`
		} `yaml:"mySql"`
	} `yaml:"db"`

	AWS struct {
		Profile string `yaml:"profile"`
		Region  string `yaml:"region"`
		Bucket  string `yaml:"bucket"`
	} `yaml:"aws"`
	ImageRootUrl     string `yaml:"imageRootUrl"`
	ImagePublicKeyID string `yaml:"imagePublicKeyID"`
	ImagePrivateKey  string `yaml:"imagePrivateKey"`
	Push             struct {
		FCM struct {
			ServiceAccountFile string `yaml:"ServiceAccountFile"`
		} `yaml:"fcm"`
	} `yaml:"push"`
	NotifyRoutine struct {
		Enabled  bool `yaml:"enabled"`
		Interval int  `yaml:"interval"`
	} `yaml:"notifyRoutine"`
}
