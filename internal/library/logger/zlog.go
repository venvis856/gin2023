package logger

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path"
	"runtime"
)

var levelMap = map[string]zapcore.Level{
	"debug":   zapcore.DebugLevel,
	"error":   zapcore.ErrorLevel,
	"warning": zapcore.WarnLevel,
	"info":    zapcore.InfoLevel,
}

type Config struct {
	File        string `help:"日志输出文件" devDefault:"$ROOT/logs/run.log" default:"$ROOT/logs/run.log"`
	FileSize    int    `help:"日志文件大小限制,单位MB" default:"100"`
	FileBackups int    `help:"最大保留日志文件数量" default:"10"`
	FileAge     int    `help:"日志文件保留天数" default:"0"`
	Level       string `help:"日志输出级别,可选[debug|info|error|warning]" default:"debug"`
	Output      string `help:"日志输出方式,可选[any|file|console]"  devDefault:"any" default:"file"`
	Encoder     string `help:"日志输出格式,可选[json|console]" default:"json"`
	Caller      bool   `help:"是否打印caller信息" default:"true"`
}

type Log struct {
	conf   *Config
	logger *zap.Logger
}

func NewZapLog(cfg *Config) *zap.Logger {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	var encoder zapcore.Encoder
	if cfg.Encoder == "console" {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	}
	//文件writeSyncer
	fileWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   cfg.File,        //日志文件存放目录
		MaxSize:    cfg.FileSize,    //文件大小限制,单位MB
		MaxBackups: cfg.FileBackups, //最大保留日志文件数量
		MaxAge:     cfg.FileAge,     //日志文件保留天数
		Compress:   false,           //是否压缩处理
	})

	level := zapcore.InfoLevel
	if _level, ok := levelMap[cfg.Level]; ok {
		level = _level
	}
	consoleOutput := zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), level)
	fileOutput := zapcore.NewCore(encoder, fileWriteSyncer, level)
	var core zapcore.Core
	switch cfg.Output {
	case "any":
		core = zapcore.NewTee(consoleOutput, fileOutput)
		break
	case "file":
		core = fileOutput
		break
	default:
		core = consoleOutput
	}
	return zap.New(core)
}

func NewLog(cfg *Config) (log *Log, err error) {
	return &Log{conf: cfg, logger: NewZapLog(cfg)}, nil
}

func (l *Log) Info(message string, fields ...zap.Field) {
	if l.conf.Caller {
		callerFields := getCallerInfoForLog()
		fields = append(fields, callerFields...)
	}
	l.logger.Info(message, fields...)
}

func (l *Log) Debug(message string, fields ...zap.Field) {
	if l.conf.Caller {
		callerFields := getCallerInfoForLog()
		fields = append(fields, callerFields...)
	}
	l.logger.Debug(message, fields...)
}

func (l *Log) Error(message string, fields ...zap.Field) {
	if l.conf.Caller {
		callerFields := getCallerInfoForLog()
		fields = append(fields, callerFields...)
	}
	l.logger.Error(message, fields...)
}

func (l *Log) Warn(message string, fields ...zap.Field) {
	if l.conf.Caller {
		callerFields := getCallerInfoForLog()
		fields = append(fields, callerFields...)
	}
	l.logger.Warn(message, fields...)
}

func getCallerInfoForLog() (callerFields []zap.Field) {
	pc, file, line, ok := runtime.Caller(2) // 回溯两层，拿到写日志的调用方的函数信息
	if !ok {
		return
	}
	funcName := runtime.FuncForPC(pc).Name()
	funcName = path.Base(funcName) //Base函数返回路径的最后一个元素，只保留函数名
	callerFields = append(callerFields, zap.String("func", funcName), zap.String("file", file), zap.Int("line", line))
	return
}
