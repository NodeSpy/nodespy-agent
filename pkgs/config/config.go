package config

import (
	"github.com/NodeSpy/nodespy-agent/pkgs/logging"
	"github.com/spf13/viper"
)

type Config struct {
}

func init() {
	viper.SetConfigName("nodespy")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("/etc/nodespy-agent/")

	viper.SetDefault("LogLevel", "Warn")
	viper.SetDefault("LogFile", "/var/log/nodespy-agent.log")

	log := logging.Configure(logging.Config{})

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// config file not found
			log.Fatal().Err(err).Msg("No config file found!")
		}
	}
}
