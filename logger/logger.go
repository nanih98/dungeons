package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

type CustomLogger struct {
	//Log *logrus.Entry
	Log *logrus.Logger
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

func (c *CustomLogger) CustomFields() *logrus.Entry {
	contextLogger := c.Log.WithFields(logrus.Fields{
		"app": "dungeons",
	})
	return contextLogger
}

func (c *CustomLogger) LogLevel(level string) {
	switch level {
	case "debug":
		c.Log.SetLevel(logrus.DebugLevel)
	case "info":
		c.Log.SetLevel(logrus.InfoLevel)
	case "warning":
		c.Log.SetLevel(logrus.WarnLevel)
	case "error":
		c.Log.SetLevel(logrus.ErrorLevel)
	case "trace":
		c.Log.SetLevel(logrus.TraceLevel)
	}
}

// Logger function
func Logger() CustomLogger {
	var log = &logrus.Logger{
		Out:       os.Stderr,
		Formatter: new(logrus.JSONFormatter),
		Hooks:     make(logrus.LevelHooks),
	}

	return CustomLogger{Log: log}
}
