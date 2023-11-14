package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger //For structured logging.

func init() {
	var err error

	// Changing predefined cofiguration for zap.logger which is valid for most of the case.
	// Instead of using EPOC time, want to use ISO 8901 formate timestamps for clarity.
	config := zap.NewProductionConfig()

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.StacktraceKey = "" // It will not set the stacktrace (while calling logger.Error())

	config.EncoderConfig = encoderConfig

	log, err = config.Build(zap.AddCallerSkip(1)) //AddCallerSkip to specify the levels to skip while logging. This means if we skip one level (which is current package dir) as we are wrapping directly, it will show exact position of invokation in the code.	
	// log, err = zap.NewProduction(zap.AddCallerSkip(1)) //Or we can use zap.NewDevlopment().

	if err != nil {
		panic(err)
	}
}

// Instead of exposing zap.Logger, preferable to expose the wrapper function.
// Helper function to wrap the zap.logger. In Future, it can be easily replaced by any other thri part logger.
//	so that we can easily change logger in future with least changes and no dependencies.
func Info(message string, fields ...zap.Field) { //Varadics --> Can pass any no of arguments; same as slice
	log.Info(message, fields...)
}

// Adding more Helper functions, creating library for popular helper functions
func Debug(message string, fields ...zap.Field) { //Varadics --> Can pass any no of arguments; same as slice
	log.Debug(message, fields...)
}

func Error(message string, fields ...zap.Field) { //Varadics --> Can pass any no of arguments; same as slice
	log.Error(message, fields...)
}
