package log

import (
	"log"

	"github.com/newrelic/go-agent"
)

func Init(appName string, appKey string, enabled bool) newrelic.Application {

	config := newrelic.NewConfig(appName, appKey)
	config.Enabled = enabled
	app, err := newrelic.NewApplication(config)

	if err != nil {
		log.Fatalf("%s", err)
	}

	return app
}
