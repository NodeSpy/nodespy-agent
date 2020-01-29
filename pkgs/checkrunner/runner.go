package checkrunner

import (
	"sync"
	"time"

	checks "github.com/NodeSpy/nodespy-agent/checks/cpu"
	"github.com/NodeSpy/nodespy-agent/pkgs/configuration"
	"github.com/NodeSpy/nodespy-agent/pkgs/logging"
)

func Run(log *logging.Logger, config *configuration.Config) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for {
			checks.Start()
			time.Sleep(time.Second * time.Duration(config.Interval))
		}
	}(&wg)

	wg.Wait()
}
