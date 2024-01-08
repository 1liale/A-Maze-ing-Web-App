package middlewares

import (
	"io"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

type CustomStdLogger struct {
	*logrus.Logger
}

func NewLogger() *CustomStdLogger {
	logger := logrus.New()

	// load env
	err := godotenv.Load()
	if err != nil {
		logger.Fatal("Error loading .env file")
	}

	// configure logger
	logLevel, err := logrus.ParseLevel(os.Getenv("LOGGER_LEVEL"))
	if err != nil {
		logLevel = logrus.InfoLevel
	}

	// log errors
	f_path := "errors.log"
	error_logrot := lumberjack.Logger{
		Filename:   f_path,
		MaxSize:    1, // MB
		MaxBackups: 3,
		MaxAge:     28, // Days
	}

	// split output to mutliple channels
	channels := io.MultiWriter(os.Stdout, &error_logrot)

	// log formatter
	logFormat := logrus.TextFormatter{
		TimestampFormat: time.RFC1123Z,
		FullTimestamp:   true,
		ForceColors:     true,
	}

	logger.SetFormatter(&logFormat)
	logger.SetLevel(logLevel)
	logger.SetOutput(channels)

	stdLogger := &CustomStdLogger{logger}

	return stdLogger
}

// Example custom error handling
var (
	invalidArgMessage      = "Invalid arg: %s"
	invalidArgValueMessage = "Invalid value for argument: %s: %v"
	missingArgMessage      = "Missing arg: %s"
)

// Example CustomStdLogger methods to print out standardized error messages
func (l *CustomStdLogger) InvalidArg(argumentName string) {
	l.Errorf(invalidArgMessage, argumentName)
}

func (l *CustomStdLogger) InvalidArgValue(argumentName string, argumentValue string) {
	l.Errorf(invalidArgValueMessage, argumentName, argumentValue)
}

func (l *CustomStdLogger) MissingArg(argumentName string) {
	l.Errorf(invalidArgValueMessage, argumentName)
}
