package libs

import (
	"fmt"
	"path/filepath"

	"github.com/jinzhu/configor"
	logger "github.com/sirupsen/logrus"
)

var Config = struct {
	Debug    bool   `default:"false" env:"Debug"`
	LogLevel string `default:"info" env:"Loglevel"`
	HTTPS    bool   `default:"false" env:"HTTPS"`
	Certpath string `default:"" env:"Certpath"`
	Certkey  string `default:"" env:"Certkey"`
	Port     int    `default:"8080" env:"PORT"`
	Host     string `default:"127.0.0.1" env:"Host"`

	DB struct {
		Prefix   string `env:"DBPrefix" default:"iris_"`
		Name     string `env:"DBName" default:"car-tms"`
		Adapter  string `env:"DBAdapter" default:"mysql"`
		Host     string `env:"DBHost" default:"127.0.0.1"`
		Port     string `env:"DBPort" default:"3306"`
		User     string `env:"DBUser" default:"root"`
		Password string `env:"DBPassword" default:"123456"`
	}
	SMS struct {
		SecretId  string `env:"SMSSecretId" default:""`
		SecretKey string `env:"SMSSecretKey" default:""`
		SdkAppid  string `env:"SMSSdkAppid" default:""`
		Sign      string `env:"SMSSign" default:""`
	}
}{}

func init() {
	configPath := filepath.Join(CWD(), "config.yml")

	if err := configor.Load(&Config, configPath); err != nil {
		logger.Println(fmt.Sprintf("Config Path:%s ,Error:%s", configPath, err.Error()))
	}
}
