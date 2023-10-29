package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logger *zap.Logger

var logLevelSeverity = map[zapcore.Level]string{
	zapcore.DebugLevel:  "DEBUG",
	zapcore.InfoLevel:   "INFO",
	zapcore.WarnLevel:   "WARNING",
	zapcore.ErrorLevel:  "ERROR",
	zapcore.DPanicLevel: "CRITICAL",
	zapcore.PanicLevel:  "ALERT",
	zapcore.FatalLevel:  "EMERGENCY",
}

func InitLogger() {
	stdout := zapcore.AddSync(os.Stdout)

	file := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "logs/app.log",
		MaxSize:    10, // megabytes
		MaxBackups: 3,
		MaxAge:     7, // days
	})

	level := zap.NewAtomicLevelAt(zap.DebugLevel)

	productionCfg := zap.NewProductionEncoderConfig()
	productionCfg.TimeKey = "timestamp"
	productionCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	productionCfg.CallerKey = "caller"
	productionCfg.EncodeCaller = zapcore.ShortCallerEncoder

	developmentCfg := zap.NewDevelopmentEncoderConfig()
	developmentCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
	developmentCfg.CallerKey = "caller"
	developmentCfg.EncodeCaller = zapcore.ShortCallerEncoder

	consoleEncoder := zapcore.NewConsoleEncoder(developmentCfg)
	fileEncoder := zapcore.NewJSONEncoder(productionCfg)

	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, stdout, level),
		zapcore.NewCore(fileEncoder, file, level),
	)

	logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	defer logger.Sync()
}

func Info(msgs ...interface{}) {
	logger.Sugar().Info(msgs...)
}

func Infof(template string, msgs ...interface{}) {
	logger.Sugar().Infof(template, msgs...)
}

func Warn(msgs ...interface{}) {
	logger.Sugar().Warn(msgs...)
}

func Warnf(template string, msgs ...interface{}) {
	logger.Sugar().Warnf(template, msgs...)
}

func Fatal(msgs ...interface{}) {
	logger.Sugar().Fatal(msgs...)
}

func Fatalf(template string, msgs ...interface{}) {
	logger.Sugar().Fatalf(template, msgs...)
}

func Panic(msgs ...interface{}) {
	logger.Sugar().Panic(msgs...)
}

func Panicf(template string, msgs ...interface{}) {
	logger.Sugar().Panicf(template, msgs...)
}

func Error(msgs ...interface{}) {
	logger.Sugar().Error(msgs...)
}

func Errorf(template string, msgs ...interface{}) {
	logger.Sugar().Errorf(template, msgs...)
}

func Debug(msgs ...interface{}) {
	logger.Sugar().Debug(msgs...)
}

func Debugf(template string, msgs ...interface{}) {
	logger.Sugar().Debugf(template, msgs...)
}
