package configuration

import (
	"github.com/NodeSpy/nodespy-agent/pkgs/logging"
	"github.com/spf13/viper"
)

type Config struct {
	APIKey       string // API Key to talk to nodespy
	Interval     int    // Seconds to wait between check runs
	DebugLogging bool   // Output debug logs
	LogFile      string // Log file location, empty for no file
}

var log *logging.Logger = logging.Configure(logging.Config{})

func init() {
	viper.SetConfigName("nodespy")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("/etc/nodespy-agent/")

	viper.SetDefault("LogLevel", "Warn")
	viper.SetDefault("LogFile", "./nodespy-agent.log")

	viper.SetDefault("Interval", 15)
}

func LoadConfig() *Config {
	var config Config
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatal().Err(err).Msg("No config file found!")
		}
	}

	err := viper.Unmarshal(&config)
	if err != nil {
		log.Fatal().Err(err).Msg("Unable to decode into struct")
	}

	return &config
}
