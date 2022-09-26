package logger

import (
	"os"

	"github.com/nanih98/dungeons/dto"
	"github.com/sirupsen/logrus"
)

type CustomLogger struct {
	Log          *logrus.Entry
	CustomFields *dto.Fields
}

func (c *CustomLogger) Info(msg string) {
	c.Log.Info(msg)
}

func (c *CustomLogger) Warn(msg string) {
	c.Log.Warn(msg)
}

func (c *CustomLogger) Fatal(msg error) {
	c.Log.Fatal(msg)
}

func (c *CustomLogger) Debug(msg string) {
	c.Log.Debug(msg)
}

func (c *CustomLogger) FuzzerFields() {
	contextLogger := c.Log.WithFields(logrus.Fields{
		"app":   c.CustomFields.App,
		"test":  c.CustomFields.Test,
		"test2": c.CustomFields.Test2,
	})
	c.Log = contextLogger
}

func (c *CustomLogger) LogLevel(level string) {
	switch level {
	case "debug":
		c.Log.Logger.SetLevel(logrus.DebugLevel)
	case "info":
		c.Log.Logger.SetLevel(logrus.InfoLevel)
	case "warning":
		c.Log.Logger.SetLevel(logrus.WarnLevel)
	case "error":
		c.Log.Logger.SetLevel(logrus.ErrorLevel)
	case "trace":
		c.Log.Logger.SetLevel(logrus.TraceLevel)
	}
}

// Logger function
func Logger() CustomLogger {
	var log = &logrus.Logger{
		Out:       os.Stderr,
		Formatter: new(logrus.JSONFormatter),
		Hooks:     make(logrus.LevelHooks),
	}

	contextLogger := log.WithFields(logrus.Fields{
		//Empty by default
	})

	return CustomLogger{Log: contextLogger, CustomFields: &dto.Fields{}}
}
