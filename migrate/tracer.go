package migrate

import (
	"go.opentelemetry.io/otel/trace"
)

// IsLoaded returns true if the tracer has been loaded.
func (m *Migrator) IsLoaded() bool {
	return m.tracer != nil
}

// Returns the wrapped tracer.
func (m *Migrator) Tracer() trace.Tracer {
	return m.tracer
}
