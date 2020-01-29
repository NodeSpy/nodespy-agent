package main

import (
	checks "github.com/NodeSpy/nodespy-agent/checks/cpu"
	"github.com/NodeSpy/nodespy-agent/pkgs/configuration"
	"github.com/NodeSpy/nodespy-agent/pkgs/logging"
)

var log *logging.Logger
var config *configuration.Config

func init() {
	config = configuration.LoadConfig()
	log = logging.Configure(logging.Config{
		Debug:   config.DebugLogging,
		LogFile: config.LogFile,
	})
}

func main() {
	checks.Start()
}
