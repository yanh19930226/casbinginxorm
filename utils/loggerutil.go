package utils

import (
	"path"

	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

// 日志名称
const (
	//日志文件名
	LOG_NAME = "system"
	//日志文件后缀
	LOG_SUFFIX = ".log"
	//单个日志文件大小，单位MB
	LOG_SIZE = 50
	//日志文件个数
	LOG_BACKUP = 10
	//日志文件最大天数
	LOG_DATE = 7
)

// 初始化日志模块
func InitLogger() {
	log_file_path := path.Join("./", LOG_NAME+LOG_SUFFIX)
	logger = logrus.New()
	setOutPut(logger, log_file_path)
	logger.SetLevel(logrus.DebugLevel)
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
}

// 设置日志输出到文件
func setOutPut(log *logrus.Logger, log_file_path string) {
	logconf := &lumberjack.Logger{
		Filename:   log_file_path,
		MaxSize:    LOG_SIZE,   // 日志文件大小，单位是 MB
		MaxBackups: LOG_BACKUP, // 最大过期日志保留个数
		MaxAge:     LOG_DATE,   // 保留过期文件最大时间，单位 天
		Compress:   true,       // 是否压缩日志，默认是不压缩。这里设置为true，压缩日志
	}
	log.SetOutput(logconf)
}

// 获取logrus操作对象
func GetLogger() *logrus.Logger {
	return logger
}
