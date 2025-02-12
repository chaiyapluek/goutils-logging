package utils

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	*zap.Logger
}

func GetLogger() *Logger {
	cfg := zap.Config{
		Level:    zap.NewAtomicLevelAt(zap.DebugLevel),
		Encoding: "json",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "time",
			LevelKey:       "level",
			MessageKey:     "message",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.FullCallerEncoder,
		},
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}
	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}

	return &Logger{
		logger,
	}
}

func (l *Logger) with(c RequestContext) *zap.Logger {
	zapfields := []zap.Field{}
	for k, v := range c.GetLogContext() {
		zapfields = append(zapfields, zap.Any(k, v))
	}
	return l.Logger.With(zapfields...)
}

func (l *Logger) Info(template string, args ...any) {
	l.Logger.Sugar().Infof(template, args...)
}

func (l *Logger) CInfo(c RequestContext, template string, args ...any) {
	l.with(c).Sugar().Infof(template, args...)
}

func (l *Logger) Warn(template string, args ...any) {
	l.Logger.Sugar().Warnf(template, args...)
}

func (l *Logger) CWarn(c RequestContext, template string, args ...any) {
	l.with(c).Sugar().Warnf(template, args...)
}

func (l *Logger) Error(template string, args ...any) {
	l.Logger.Sugar().Errorf(template, args...)
}

func (l *Logger) CError(c RequestContext, template string, args ...any) {
	l.with(c).Sugar().Errorf(template, args...)
}

func (l *Logger) Debug(template string, args ...any) {
	l.Logger.Sugar().Debugf(template, args...)
}

func (l *Logger) CDebug(c RequestContext, template string, args ...any) {
	l.with(c).Sugar().Debugf(template, args...)
}
