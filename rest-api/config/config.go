package config

import (
    "errors"
    log "github.com/sirupsen/logrus"
    "github.com/spf13/viper"
)

var (
	STAGES = []string{"default", "local", "dev", "prd", "test"}
)

func init(){
    Config =&config{}
    viper.AddConfigPath(".")
    viper.AddConfigPath("config")
    viper.SetEnvPrefix("EVERYBODY")
    viper.AutomaticEnv()

    found := false

    for _, stage := range STAGES{
        viper.SetConfigName(stage)
        if err := viper.ReadInConfig(); err != nil{
            if !errors.As(err, &viper.ConfigFileNotFoundError{}){
                log.Fatal(err)
            }
        } else{
            found = true
            log.Infof("%s stage에 대한 설정파일을 발견했습니다.", stage)

            if err := viper.Unmarshal(Config); err != nil{
                log.Fatal(err)
            }
        }
    }

    if !found {
        log.Fatalf("설정파일을 하나도 찾지 못했습니다. 다음 stage 중 하나에 대한 설정파일을 생성해주세요. %#v", STAGES)
    }
    log.Info(Config)
}

var Config *config

type config struct{
    DB struct{
        MySQL struct{
            Host         string `yaml:"host"`
			Port         int    `yaml:"port"`
			DatabaseName string `yaml:"databaseName"`
			User         string `yaml:"user"`
			Password     string `yaml:"password"`
        } `yaml:"mySql"`
    } `yaml:"db"`
}