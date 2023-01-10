package logger

//import (
//	"fmt"
//	"go.uber.org/zap"
//	"go.uber.org/zap/zapcore"
//	"gopkg.in/natefinch/lumberjack.v2"
//	gormLogger "gorm.io/gorm/logger"
//	"os"
//)
//
//type LogType int    // 日志类型
//type LogOutType int // 日志输出类型
//
//const (
//	LOG_TYPE_DEBUG LogType = 1 // 调试日志
//	LOG_TYPE_ERROR LogType = 2 // 错误日志
//	LOG_TYPE_DB    LogType = 3 // 数据库日志
//
//	LOG_OUT_ALL    LogOutType = 1 // 同时输出控制台和文件
//	LOG_OUT_STDOUT LogOutType = 2 // 控制台输出 默认
//	LOG_OUT_FILE   LogOutType = 3 // 文件输出
//)
//
//type LogConfig struct {
//	LogType    LogType
//	LogOut     LogOutType
//	LogLvl     zapcore.LevelEnabler
//	Path       string // 日志文件的位置，也就是路径
//	MaxSize    int    // 在进行切割之前，日志文件的最大大小（以MB为单位）
//	MaxBackups int    // 保留旧文件的最大个数
//	MaxAge     int    // 保留旧文件的最大天数
//	Compress   bool   // 是否压缩/归档旧文件
//}
//
//func (cfg *LogConfig) GetLumberJackCfg() *lumberjack.Logger {
//	var outFile string
//	if cfg.LogType == LOG_TYPE_DEBUG {
//		outFile = fmt.Sprintf("%s/info.log", cfg.Path)
//	} else if cfg.LogType == LOG_TYPE_ERROR {
//		outFile = fmt.Sprintf("%s/err.log", cfg.Path)
//	} else if cfg.LogType == LOG_TYPE_DB {
//		outFile = fmt.Sprintf("%s/gorm.log", cfg.Path)
//	}
//
//	msg := &lumberjack.Logger{
//		Filename:   outFile,
//		MaxSize:    cfg.MaxSize,
//		MaxBackups: cfg.MaxBackups,
//		MaxAge:     cfg.MaxAge,
//		Compress:   cfg.Compress,
//	}
//	return msg
//}
//
//func (cfg *LogConfig) GetLogWriter() zapcore.WriteSyncer {
//	if cfg.LogOut == LOG_OUT_ALL {
//		std := zapcore.AddSync(os.Stdout)
//		stdWriteSyncer := zapcore.AddSync(std)
//
//		file := zapcore.AddSync(cfg.GetLumberJackCfg())
//		jackWriteSyncer := zapcore.AddSync(file)
//		return zapcore.NewMultiWriteSyncer(stdWriteSyncer, jackWriteSyncer)
//	} else if cfg.LogOut == LOG_OUT_FILE {
//		return zapcore.AddSync(cfg.GetLumberJackCfg())
//	}
//	return os.Stdout
//}
//
//func (cfg *LogConfig) GetMultiWriter() zapcore.WriteSyncer {
//	std := zapcore.AddSync(os.Stdout)
//	stdWriteSyncer := zapcore.AddSync(std)
//	c := cfg.GetLumberJackCfg()
//	c.Filename = fmt.Sprintf("%s/err.log", cfg.Path)
//	file := zapcore.AddSync(c)
//	jackWriteSyncer := zapcore.AddSync(file)
//	return zapcore.NewMultiWriteSyncer(stdWriteSyncer, jackWriteSyncer)
//}
//
//func NewLogger(cfg LogConfig) *zap.Logger {
//	encoderConfig := zap.NewProductionEncoderConfig()
//	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
//	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
//	encoder := zapcore.NewConsoleEncoder(encoderConfig)
//
//	writerSyncer := cfg.GetLogWriter()
//
//	var loglvl zapcore.LevelEnabler
//	if cfg.LogLvl != nil {
//		loglvl = cfg.LogLvl
//	} else {
//		loglvl = zapcore.DebugLevel
//	}
//	core := zapcore.NewCore(encoder, writerSyncer, loglvl)
//	opt := zap.ErrorOutput(cfg.GetMultiWriter())
//	//opt := zap.AddCaller()
//	logger := zap.New(core, opt)
//	//defer logger.Sync() // flushes buffer, if any
//	return logger
//}
//
//func NewDefaultLogger(typ LogType, out LogOutType, name string) *zap.SugaredLogger {
//	cfg := LogConfig{
//		LogType:    typ,
//		LogOut:     out,
//		Path:       "log",
//		MaxSize:    10,
//		MaxBackups: 50,
//		MaxAge:     30,
//		Compress:   true,
//	}
//	l := NewLogger(cfg)
//	return l.Named(name).Sugar()
//}
//
//func NewDefaultGormLogger(out LogOutType, gormCfg gormLogger.Config) gormLogger.Interface {
//	cfg := LogConfig{
//		LogType:    LOG_TYPE_DB,
//		LogOut:     out,
//		Path:       "log",
//		MaxSize:    10,
//		MaxBackups: 50,
//		MaxAge:     30,
//		Compress:   true,
//	}
//	l := NewLogger(cfg)
//	writer := &GormWriter{
//		Logger: l.Sugar(),
//	}
//	return gormLogger.New(writer, gormCfg)
//}
//
//type GormWriter struct {
//	Logger *zap.SugaredLogger
//}
//
//func (w *GormWriter) Printf(msg string, data ...interface{}) {
//	//fmt.Println("--->", msg)
//	//for i, datum := range data {
//	//	fmt.Println("--->", i, datum)
//	//}
//	//fmt.Printf(msg, data...)
//	w.Logger.Infof(msg, data...)
//}
