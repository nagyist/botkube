package audit

import (
	"context"

	"github.com/sirupsen/logrus"
)

var _ AuditReporter = (*NoopAuditReporter)(nil)

// NoopAuditReporter is the NOOP audit reporter
type NoopAuditReporter struct {
	log logrus.FieldLogger
}

func newNoopAuditReporter(logger logrus.FieldLogger) *NoopAuditReporter {
	return &NoopAuditReporter{
		log: logger,
	}
}

// ReportExecutorAuditEvent is a NOOP
func (r *NoopAuditReporter) ReportExecutorAuditEvent(ctx context.Context, e ExecutorAuditEvent) error {
	return nil
}

// ReportSourceAuditEvent is a NOOP
func (r *NoopAuditReporter) ReportSourceAuditEvent(ctx context.Context, e SourceAuditEvent) error {
	return nil
}
