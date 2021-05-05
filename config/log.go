package config

import (
	"github.com/natefinch/lumberjack"
	"go-server/global"
	"go-server/initialize"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func getEncoder() zapcore.Encoder {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder, // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,    // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder, // 全路径编码器
	}
	return zapcore.NewConsoleEncoder(encoderConfig)
}
func getLogWriter(logPath string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   logPath, // 日志文件路径，默认 os.TempDir()
		MaxSize:    5,       // 每个日志文件保存10M，默认 100M
		MaxBackups: 90,      // 保留90个备份，默认不限
		MaxAge:     90,      // 保留90天，默认不限
		Compress:   true,    // 是否压缩，默认不压缩
	}
	return zapcore.AddSync(lumberJackLogger)
}

func initLog() {

	write := getLogWriter(initialize.Application.LogPath)
	core := zapcore.NewTee(zapcore.NewCore(getEncoder(), zapcore.AddSync(write), zap.InfoLevel))

	// 构建日志
	logger := zap.New(core, zap.AddCaller())
	logger.Info("log 初始化成功")
	global.Logger = logger

}
