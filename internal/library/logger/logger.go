package logger

import (
	"github.com/gogf/gf/util/gconv"
	"github.com/golang-module/carbon"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"path/filepath"
	"sync"
)

type Logger struct {
	mu      sync.Mutex
	conf    Config
	loggers map[string]*zap.Logger
}

func NewLogger(conf Config) (logger *Logger, err error) {
	return &Logger{
		conf:    conf,
		loggers: make(map[string]*zap.Logger),
	}, nil
}

/*
*
日志每天的形式
*/
func (l *Logger) Daily(filepath string, level string, data interface{}, datas ...interface{}) {
	//日期目录
	fileKey := filepath + "_" + level + "_" + carbon.Now().Format("Ymd") + ".log"
	msg := gconv.String(data)
	if len(datas) > 0 {
		for i := 0; i < len(datas); i++ {
			msg += " | " + gconv.String(datas[i])
		}
	}
	var zLevel zapcore.Level
	var ok bool
	if zLevel, ok = levelMap[level]; !ok {
		zLevel = zapcore.DebugLevel
	}
	l.getLogger(fileKey).Log(zLevel, msg)
}

func (l *Logger) Write(filepath string, level string, data interface{}, datas ...interface{}) {
	fileKey := filepath + "_" + level  + ".log"
	msg := gconv.String(data)
	if len(datas) > 0 {
		for i := 0; i < len(datas); i++ {
			msg += " | " + gconv.String(datas[i])
		}
	}
	var zLevel zapcore.Level
	var ok bool
	if zLevel, ok = levelMap[level]; !ok {
		zLevel = zapcore.DebugLevel
	}
	l.getLogger(fileKey).Log(zLevel, msg)
}

func (l *Logger) getLogger(fileKey string) *zap.Logger {
	l.mu.Lock()
	defer l.mu.Unlock()
	if v, ok := l.loggers[fileKey]; ok {
		return v
	}
	conf := l.conf
	conf.File = filepath.Join(filepath.Dir(conf.File), fileKey)
	_zlog := NewZapLog(&conf)
	l.loggers[fileKey] = _zlog
	return _zlog
}
