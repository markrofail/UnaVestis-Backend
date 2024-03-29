package config

import (
	"fmt"
	"github.com/spf13/viper"
)

// Config is global object that holds all application level variables.
var Config appConfig

type appConfig struct {
	// the server port. Defaults to 8080
	ServerPort int `mapstructure:"server_port"`
}

// LoadConfig loads config from files
func LoadConfig(configPaths ...string) error {
	v := viper.New()
	v.SetConfigName("example")
	v.SetConfigType("yaml")
	v.SetEnvPrefix("unaVestis")
	v.AutomaticEnv()

	//Config.DSN = v.Get("DSN").(string)
	//Config.ApiKey = v.Get("API_KEY").(string)
	v.SetDefault("server_port", 1234)

	for _, path := range configPaths {
		v.AddConfigPath(path)
	}
	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to read the configuration file: %s", err)
	}
	return v.Unmarshal(&Config)
}
