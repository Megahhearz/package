package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger — интерфейс логгера с методами для информационных и ошибок логов.
type Logger interface {
	Info(layer string, method string, msg string, args ...interface{})
	Error(layer string, method string, msg string, err error, args ...interface{})
}

// logger — структура-обертка над slog.Logger.
type logger struct {
	logger *zap.Logger
}

// New создает новый экземпляр структурированного логгера,
// выводящего сообщения в stdout с уровнем Info и выше.
func New() *logger {
	cfg := zap.Config{
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		Level:            zap.NewAtomicLevelAt(zapcore.InfoLevel),
	}
	zapLogger, err := cfg.Build()

	if err != nil {
		panic("failed to build Logger" + err.Error())
	}

	return &logger{
		logger: zapLogger,
	}
}

// Info логирует информационное сообщение с указанием слоя, метода и дополнительных деталей.
func (l *logger) Info(layer string, method string, msg string, args ...interface{}) {
	l.logger.Info(
		msg,
		zap.String("layer", layer),
		zap.String("method", method),
		zap.Any("details", args), // Дополнительные параметры логируются как details.
	)
}

// Error логирует ошибку с указанием слоя, метода, сообщения, самой ошибки и дополнительных деталей.
func (l *logger) Error(layer string, method string, msg string, err error, args ...interface{}) {
	l.logger.Error(
		msg,
		zap.String("layer", layer),
		zap.String("method", method),
		zap.Any("error", err),    // Ошибка логируется в поле "error".
		zap.Any("details", args), // Дополнительные параметры — в details.
	)
}
