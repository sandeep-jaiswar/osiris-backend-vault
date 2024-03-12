package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	Log      *zap.Logger
	filename = "session_Service.log"
)

func NewConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		MessageKey:       "message",
		LevelKey:         "level",
		TimeKey:          "time",
		NameKey:          "logger",
		CallerKey:        "caller",
		FunctionKey:      "function",
		StacktraceKey:    "stacktrace",
		SkipLineEnding:   false,
		LineEnding:       zapcore.DefaultLineEnding,
		EncodeLevel:      zapcore.LowercaseLevelEncoder,
		EncodeTime:       zapcore.ISO8601TimeEncoder,
		EncodeDuration:   zapcore.SecondsDurationEncoder,
		EncodeCaller:     zapcore.ShortCallerEncoder,
		EncodeName:       zapcore.FullNameEncoder,
		ConsoleSeparator: " | ",
	}
}

func EnableLogger() {
	config := NewConfig()
	fileEncoder := zapcore.NewJSONEncoder(config)
	consoleEncoder := zapcore.NewJSONEncoder(config)
	logFile, _ := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	writer := zapcore.AddSync(logFile)
	defaultLogLevel := zapcore.DebugLevel
	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, writer, defaultLogLevel),
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), defaultLogLevel),
	)
	Log = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	Log.Info("Logger enabled successfully")
}
