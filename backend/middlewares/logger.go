package middlewares

import (
	// need it to print out errors before logger is properly setup
	"io"
	"os"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

type CustomStdLogger struct {
	*logrus.Logger
}

type LogWriterHook struct {
	Writer    io.Writer
	LogLevels []logrus.Level
}

func (hook *LogWriterHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		return err
	}
	_, err = hook.Writer.Write([]byte(line))
	return err
}

func (hook *LogWriterHook) Levels() []logrus.Level {
	return hook.LogLevels
}

// configure logger output level
func ParseLogLevels(env_name string) []logrus.Level {
	log_levels_env := os.Getenv(env_name)
	if log_levels_env == "" {
		return nil
	}
	log_level_strings := strings.Split(log_levels_env, ",")
	log_level_arr := make([]logrus.Level, len(log_level_strings))
	for i, log_level := range log_level_strings {
		parsed_level, err := logrus.ParseLevel(log_level)
		if err != nil {
			logrus.Fatal(err)
		}
		log_level_arr[i] = parsed_level
	}

	return log_level_arr
}

func InitLogger() *CustomStdLogger {
	// set logger configurations
	logger := logrus.New()
	log_format := logrus.TextFormatter{
		TimestampFormat: time.RFC1123Z,
		FullTimestamp:   true,
		ForceColors:     true,
	}
	logger.SetOutput(io.Discard)
	logger.SetFormatter(&log_format)

	// create api.log if it doesn't exist yet
	// log_path := "api.log"
	// _, err := os.OpenFile(log_path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// if err != nil {
	// 	logrus.Fatal(err)
	// }

	// file_log_levels := ParseLogLevels("FILE_LOG_LEVELS")
	output_log_levels := ParseLogLevels("OUTPUT_LOG_LEVELS")

	// // setup logrotation using lumberjack
	// error_logrot := lumberjack.Logger{
	// 	Filename:   log_path,
	// 	MaxSize:    1, // MB
	// 	MaxBackups: 3,
	// 	MaxAge:     28, // Days
	// }

	// logger.AddHook(&LogWriterHook{ // Send logs with level higher than warning to stderr
	// 	Writer:    &error_logrot,
	// 	LogLevels: file_log_levels,
	// })
	logger.AddHook(&LogWriterHook{ // Send info and debug logs to stdout
		Writer:    os.Stdout,
		LogLevels: output_log_levels,
	})

	return &CustomStdLogger{logger}
}
