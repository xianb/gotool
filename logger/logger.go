package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

type LogOut int // 日志输出类型

const (
	LOG_OUT_ALL    LogOut = 1 // 同时输出控制台和文件
	LOG_OUT_STDOUT LogOut = 2 // 控制台输出 默认
	LOG_OUT_FILE   LogOut = 3 // 文件输出
)

type Logger struct {
	DebugLogger *zap.SugaredLogger // 只用于控制台输出
	InfoLogger  *zap.SugaredLogger // 控制台与文件结合输出
	ErrLogger   *zap.SugaredLogger // 控制台与文件结合输出
}

func GetDefaultJackCfg(name string) *lumberjack.Logger {
	return &lumberjack.Logger{
		Filename:   fmt.Sprintf("log/%s.log", name),
		MaxSize:    10,
		MaxBackups: 50,
		MaxAge:     30,
		Compress:   true,
	}
}

func NewLogger(out LogOut, name string) *Logger {
	return &Logger{
		DebugLogger: NewSugaredLogger(name, LOG_OUT_STDOUT, ""),
		InfoLogger:  NewSugaredLogger(name, out, "info"),
		ErrLogger:   NewSugaredLogger(name, out, "err"),
	}
}

func NewSugaredLogger(name string, out LogOut, fileName string) *zap.SugaredLogger {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	writer := GetLogWriter(out, fileName)
	core := zapcore.NewCore(encoder, writer, zapcore.DebugLevel)
	opt := zap.AddCaller() // 打印调用行数的行号和文件名
	logger := zap.New(core, opt)
	//logger := zap.New(core)
	return logger.Named(name).Sugar()
}

func GetLogWriter(out LogOut, fileName string) zapcore.WriteSyncer {
	if out == LOG_OUT_ALL {
		std := zapcore.AddSync(os.Stdout)
		stdWriteSyncer := zapcore.AddSync(std)
		file := zapcore.AddSync(GetDefaultJackCfg(fileName))
		jackWriteSyncer := zapcore.AddSync(file)
		return zapcore.NewMultiWriteSyncer(stdWriteSyncer, jackWriteSyncer)
	} else if out == LOG_OUT_FILE {
		return zapcore.AddSync(GetDefaultJackCfg(fileName))
	}
	return os.Stdout
}

func (l *Logger) Debug(args ...interface{}) {
	l.DebugLogger.Debug(args...)
}
func (l *Logger) Debugf(template string, args ...interface{}) {
	l.DebugLogger.Debugf(template, args...)
}
func (l *Logger) Info(args ...interface{}) {
	l.InfoLogger.Info(args...)
}
func (l *Logger) Infof(template string, args ...interface{}) {
	l.InfoLogger.Infof(template, args...)
}
func (l *Logger) Error(args ...interface{}) {
	l.ErrLogger.Error(args...)
}
func (l *Logger) Errorf(template string, args ...interface{}) {
	l.ErrLogger.Errorf(template, args...)
}
