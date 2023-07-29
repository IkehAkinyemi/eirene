package configs

import (
	"github.com/spf13/viper"
)

// A Configs defines the expected config values.
type Configs struct {
	Env           string `mapstructure:"ENVIRONMENT"`
	ServerAddress string `mapstructure:"HTTP_SERVER_ADDRESS"`
}

// ParseConfigs parses the configuration files.
func ParseConfigs(path string) (configs Configs, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("secrets")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&configs)
	return
}
