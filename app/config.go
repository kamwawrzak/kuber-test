package app

import (
	"log"
	"os"

	yaml "gopkg.in/yaml.v2"
	"github.com/kelseyhightower/envconfig"
)

const (
	configFileEnvVar = "SERVICE_CONFIG_FILE"
	defaultFileName = "service-config.yaml"
)

type configWrapper struct {
	Service Config `yaml:"service" envconfig:"service"`
}

type Config struct{
	Port int `yaml:"port"`
	Rdbs	RdbsConfig

}

type RdbsConfig struct {
	User string `validate:"nonzero"`
	Password string `validate:"nonzero"`
	Port int `validate:"nonzero"`
}

func getConfigfileName() (string, string) {
	if fromEnv, ok := os.LookupEnv(configFileEnvVar); ok{
		msg := "Using config file from env vars"
		return fromEnv, msg
	}
	msg := "Using default config file"
	return defaultFileName, msg
}

func NewConfig() (Config, string) {
	wrapper := configWrapper{}
	configFile, msg := getConfigfileName()

	yfile, err := os.ReadFile(configFile)
	if err != nil {
		log.Fatal(msg, err)
	}

	if err = yaml.Unmarshal(yfile, &wrapper); err != nil {
		log.Fatal(msg, err)
	}

	if err = envconfig.Process("", &wrapper); err != nil {
		log.Fatal(msg, err)
	}

	return wrapper.Service, ""
}