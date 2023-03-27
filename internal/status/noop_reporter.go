package status

import (
	"context"
)

var _ StatusReporter = (*NoopStatusReporter)(nil)

type NoopStatusReporter struct{}

func (n NoopStatusReporter) ReportDeploymentStartup(context.Context) error {
	return nil
}

func (n NoopStatusReporter) ReportDeploymentShutdown(context.Context) error {
	return nil
}

func (n NoopStatusReporter) ReportDeploymentFailed(context.Context) error {
	return nil
}

func (n NoopStatusReporter) SetResourceVersion(int) {
}

func newNoopStatusReporter() *NoopStatusReporter {
	return &NoopStatusReporter{}
}