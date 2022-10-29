package tracer

import (
	"fmt"
	"os"

	"github.com/felixa1996/go_next_be/app/config"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func NewRelicTracer(config config.Config) *newrelic.Application {
	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName(config.AppName),
		newrelic.ConfigLicense(config.NewrelicLicenseKey),
		newrelic.ConfigDebugLogger(os.Stdout),
	)
	if nil != err {
		fmt.Println(err)
		os.Exit(1)
	}
	return app
}
