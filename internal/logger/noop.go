package logger

import (
	"io"
	//nolint:all
	uberzap "go.uber.org/zap"

	"sigs.k8s.io/controller-runtime/pkg/log/zap"
)

// Noop discards all logged messages.
func Noop() *uberzap.Logger {
	return zap.NewRaw(zap.WriteTo(io.Discard))
}
