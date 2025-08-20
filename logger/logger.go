package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger interface {
	Info(layer string, method string, msg string, args ...interface{})
	Error(layer string, method string, msg string, err error, args ...interface{})
}

type logger struct {
	logger *zap.Logger
}

func New(loglevel string) (Logger, error) {
	config := zap.Config{
		Level:            getZapLevel(loglevel),
		Encoding:         "json",
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:   "message",
			LevelKey:     "level",
			TimeKey:      "time",
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	zapLogger, err := config.Build()
	if err != nil {
		return nil, err
	}

	return &logger{
		logger: zapLogger,
	}, nil
}

func getZapLevel(logLevel string) zap.AtomicLevel {
	switch logLevel {
	case "debug":
		return zap.NewAtomicLevelAt(zap.DebugLevel)
	case "info":
		return zap.NewAtomicLevelAt(zap.InfoLevel)
	case "warn":
		return zap.NewAtomicLevelAt(zap.WarnLevel)
	case "error":
		return zap.NewAtomicLevelAt(zap.ErrorLevel)
	default:
		return zap.NewAtomicLevelAt(zap.InfoLevel)
	}
}

func (l *logger) Debug(layer string, method string, msg string, args ...interface{}) {
	l.logger.Debug(
		msg,
		zap.String("layer", layer),
		zap.String("method", method),
		zap.Any("details", args),
	)
}

func (l *logger) Info(layer string, method string, msg string, args ...interface{}) {
	l.logger.Info(
		msg,
		zap.String("layer", layer),
		zap.String("method", method),
		zap.Any("details", args),
	)
}

func (l *logger) Error(layer string, method string, msg string, err error, args ...interface{}) {
	l.logger.Error(
		msg,
		zap.String("layer", layer),
		zap.String("method", method),
		zap.Error(err),
		zap.Any("details", args),
	)
}

func (l *logger) Warn(layer string, method string, msg string, err error, args ...interface{}) {
	if err != nil {
		l.logger.Warn(
			msg,
			zap.String("layer", layer),
			zap.String("method", method),
			zap.Error(err),
			zap.Any("details", args),
		)
	} else {
		l.logger.Warn(
			msg,
			zap.String("layer", layer),
			zap.String("method", method),
			zap.Any("details", args),
		)
	}
}
