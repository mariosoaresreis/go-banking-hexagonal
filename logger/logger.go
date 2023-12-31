package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func init() {
	var err error
	config := zap.NewProductionConfig()
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig = encoderConfig

	log, err = config.Build(zap.AddCallerSkip(1))

	if err != nil {
		panic(err)
	}
}

func Info(mensagem string, fields ...zap.Field) {
	log.Info(mensagem, fields...)
}

func Error(mensagem string, fields ...zap.Field) {
	log.Error(mensagem, fields...)
}

func Fatal(mensagem string, fields ...zap.Field) {
	log.Fatal(mensagem, fields...)
}
