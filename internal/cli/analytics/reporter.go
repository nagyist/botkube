package analytics

import (
	"fmt"

	"github.com/spf13/cobra"
	"k8s.io/utils/strings"

	"github.com/kubeshop/botkube/internal/cli"
)

const (
	defaultCliVersion    = "v9.99.9-dev"
	printAPIKeyCharCount = 3
)

var (
	// APIKey contains the API key for external analytics service. It is set during application build.
	APIKey string = ""
)

// Reporter defines behavior for reporting analytics.
type Reporter interface {
	ReportCommand(cmd string) error
	ReportError(err error, cmd string) error
	Close()
}

// NewReporter creates a new Reporter instance.
func GetReporter(cmd cobra.Command) Reporter {
	if APIKey == "" {
		printWhenVerboseEnabled(cmd, "Telemetry is disabled as the API key wasn't provided.")
		return &NoopReporter{}
	}

	conf := cli.NewConfig()
	if conf.IsTelemetryDisabled() {
		printWhenVerboseEnabled(cmd, "Telemetry is disabled based on config value.")
		return &NoopReporter{}
	}

	// Create segment reporter if telemetry enabled and API key is set
	r, err := NewSegmentReporter(APIKey)
	if err != nil {
		// do not crash on telemetry errors
		printWhenVerboseEnabled(cmd, "Telemetry is disabled due to reporter misconfiguration.")
		return &NoopReporter{}
	}

	printWhenVerboseEnabled(cmd, fmt.Sprintf("Telemetry is enabled. Using API Key starting with %q...", strings.ShortenString(APIKey, printAPIKeyCharCount)))
	return r
}

func printWhenVerboseEnabled(cmd cobra.Command, s string) {
	if cli.VerboseMode.IsEnabled() {
		cmd.Println(s)
	}
}
