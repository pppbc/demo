package logs

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Loggers *zap.Logger

// logpath 日志文件路径
// loglevel 日志级别
func InitLogger(logpath string, loglevel string) *zap.Logger {

	hook := lumberjack.Logger{
		Filename:   logpath, // 日志文件路径
		MaxSize:    1024,    // megabytes
		MaxBackups: 3,       // 最多保留3个备份
		MaxAge:     7,       //days
		Compress:   true,    // 是否压缩 disabled by default
	}
	w := zapcore.AddSync(&hook)

	var level zapcore.Level
	switch loglevel {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "error":
		level = zap.ErrorLevel
	default:
		level = zap.InfoLevel
	}
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		w,
		level,
	)

	logger := zap.New(core)
	logger.Info("DefaultLogger init success")

	return logger

}

func init() {
	//日志
	Loggers = InitLogger("log_info.log", "info")
}
