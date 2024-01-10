package middlewares

import (
	"io"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
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

func NewCustomLogger() *CustomStdLogger {
	// set logger configurations
	logger := logrus.New()
	logFormat := logrus.TextFormatter{
		TimestampFormat: time.RFC1123Z,
		FullTimestamp:   true,
		ForceColors:     true,
	}
	logger.SetOutput(io.Discard)
	logger.SetFormatter(&logFormat)

	// load env
	err := godotenv.Load()
	if err != nil {
		logger.Fatal("Error loading .env file")
	}

	// create api.log if it doesn't exist yet
	log_path := "api.log"
	_, err = os.OpenFile(log_path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		logger.Errorf("Error opening %s file", log_path)
	}

	// configure logger output level
	LOG_LEVELS := os.Getenv("LOG_LEVELS")
	logLevelStrings := strings.Split(LOG_LEVELS, ",")
	logLevelArr := make([]logrus.Level, len(logLevelStrings))
	for i, logLevel := range logLevelStrings {
		logLevelArr[i], err = logrus.ParseLevel(logLevel)
	}

	// setup logrotation using lumberjack
	error_logrot := lumberjack.Logger{
		Filename:   log_path,
		MaxSize:    1, // MB
		MaxBackups: 3,
		MaxAge:     28, // Days
	}

	logger.AddHook(&LogWriterHook{ // Send logs with level higher than warning to stderr
		Writer:    &error_logrot,
		LogLevels: logLevelArr,
	})
	logger.AddHook(&LogWriterHook{ // Send info and debug logs to stdout
		Writer: os.Stdout,
		LogLevels: []logrus.Level{
			logrus.InfoLevel,
			logrus.DebugLevel,
		},
	})

	return &CustomStdLogger{logger}
}

var (
	invalidArgMessage      = "Invalid arg: %s"
	invalidArgValueMessage = "Invalid value for argument: %s: %v"
	missingArgMessage      = "Missing arg: %s"
)

// CustomStdLogger methods to print out standardized error messages
func (l *CustomStdLogger) InvalidArg(argumentName string) {
	l.Errorf(invalidArgMessage, argumentName)
}

func (l *CustomStdLogger) InvalidArgValue(argumentName string, argumentValue string) {
	l.Errorf(invalidArgValueMessage, argumentName, argumentValue)
}

func (l *CustomStdLogger) MissingArg(argumentName string) {
	l.Errorf(invalidArgValueMessage, argumentName)
}
