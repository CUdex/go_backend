package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"path/filepath"
	"strings"
)

var logger *zap.Logger
var logFile *os.File

func init() {
	var err error
	// 파일 오픈
	logFile, err = os.OpenFile("/var/log/tag.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		panic(err)
	}

	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	config = zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.CallerEncoder(func(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
			fileName := strings.Split(filepath.Base(caller.FullPath()), ":")[0]  
			enc.AppendString(fileName)
		}),
	}
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(config),
		zapcore.AddSync(logFile),
		zapcore.InfoLevel,
	)

	logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel), zap.AddCallerSkip(1))

	if err != nil {
		panic(err)
	}
}

// Info 로그 메시지를 기록합니다.
func Info(message string, fields ...zap.Field) {
	if logger != nil {
		logger.Info(message, fields...)
	} else {
		panic("logger is not initialized")
	}
}

// Error 로그 메시지를 기록합니다.
func Error(message string, fields ...zap.Field) {
	if logger != nil {
		logger.Error(message, fields...)
	} else {
		panic("logger is not initialized")
	}
}

// 프로그램 종료 시 파일 close 및 logger 메모리 정리
func Close() {
	if logFile != nil {
		logFile.Close()
		logger.Sync()
	}
}
