package utils

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type LogConfig struct {
	Level    string   `mapstructure:"level"`
	LogLevel LogLevel `mapstructure:"-" yaml:"-"`
	File     string   `mapstructure:"file"`
}

type MangoLogger struct {
	LogConfig *LogConfig
	LogError  *log.Logger
	LogWarn   *log.Logger
	LogInfo   *log.Logger
	LogDebug  *log.Logger
}

var MangoLog MangoLogger

func (ml *MangoLogger) Setup(conf LogConfig) {
	file := os.Stdout
	errFile := os.Stderr
	var err error
	if conf.File != "" {
		file, err = os.OpenFile(conf.File, os.O_APPEND, os.ModeAppend)
		checkFatal(err)
		errFile, err = os.OpenFile(conf.File, os.O_APPEND, os.ModeAppend)
		checkFatal(err)
	}
	conf.LogLevel = logLevelParse(conf.Level)
	flags := log.LstdFlags | log.LUTC | log.Llongfile | log.Lmsgprefix

	ml.LogConfig = &conf

	ml.LogDebug = log.New(file, strings.ToUpper(LogLevelDebug.String())+": ", flags)
	ml.LogInfo = log.New(file, strings.ToUpper(LogLevelInfo.String())+": ", flags)
	ml.LogWarn = log.New(file, strings.ToUpper(LogLevelWarn.String())+": ", flags)
	ml.LogError = log.New(errFile, strings.ToUpper(LogLevelError.String())+": ", flags)
}

func (ml *MangoLogger) Debug(v ...any) {
	if ml.LogConfig.LogLevel <= LogLevelDebug {
		if ml.LogDebug.Writer() == io.Discard {
			return
		}
		ml.LogDebug.Output(2, fmt.Sprintln(v...))
	}
}

func (ml *MangoLogger) Info(v ...any) {
	if ml.LogConfig.LogLevel <= LogLevelInfo {
		if ml.LogInfo.Writer() == io.Discard {
			return
		}
		ml.LogInfo.Output(2, fmt.Sprintln(v...))
	}
}

func (ml *MangoLogger) Warn(v ...any) {
	if ml.LogConfig.LogLevel <= LogLevelWarn {
		if ml.LogWarn.Writer() == io.Discard {
			return
		}
		ml.LogWarn.Output(2, fmt.Sprintln(v...))
	}
}

func (ml *MangoLogger) Error(v ...any) {
	if ml.LogConfig.LogLevel <= LogLevelError {
		if ml.LogError.Writer() == io.Discard {
			return
		}
		ml.LogError.Output(2, fmt.Sprintln(v...))
	}
}

func (ml *MangoLogger) Fatal(v ...any) {
	if ml.LogConfig.LogLevel <= LogLevelDebug {
		panic(v)
	}
	if ml.LogConfig.LogLevel <= LogLevelError {
		ml.LogError.Output(2, fmt.Sprintln(v...))
		os.Exit(1)
	}
}

func (ml *MangoLogger) fatal(v ...any) {
	if ml.LogConfig.LogLevel <= LogLevelDebug {
		panic(v)
	}
	if ml.LogConfig.LogLevel <= LogLevelError {
		ml.LogError.Output(3, fmt.Sprintln(v...))
	}
}

type LogLevel int

const (
	LogLevelDebug LogLevel = iota
	LogLevelInfo
	LogLevelWarn
	LogLevelError
)

func (l LogLevel) String() string {
	return [...]string{"Debug", "Info", "Warn", "Error"}[l]
}

func logLevelParse(val string) LogLevel {
	switch strings.ToLower(val) {
	case "debug":
		return LogLevelDebug
	case "info":
		return LogLevelInfo
	case "warn":
		return LogLevelWarn
	case "error":
		return LogLevelError
	default:
		return LogLevelInfo
	}
}
