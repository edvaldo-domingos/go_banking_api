package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)


var log *zap.Logger

func init() {
	var err error
	log, err = zap.NewProduction(zap.AddCallerSkip(1))
	config := zap.NewProductionConfig()

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig =encoderConfig

	log , err = config.Build(zap.AddCallerSkip(1))

	if err != nil {
		panic(err)
	}

}

func Info(message string, fields ...zap.Field){
	log.Info(message, fields...)
}

func Debug(message string, fields ...zap.Field){
	log.Info(message, fields...)
}

func Error(message string, fields ...zap.Field){
	log.Info(message, fields...)
}