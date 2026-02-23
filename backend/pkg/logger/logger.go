package logger

import (
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

type Config struct {
	Level         string `default:"debug"         envconfig:"LOGGER_LEVEL"`
	PrettyConsole bool   `default:"false"         envconfig:"LOGGER_PRETTY_CONSOLE"`
}

func Init(c Config) {
	logrus.SetOutput(os.Stdout)

	level, err := logrus.ParseLevel(strings.ToLower(c.Level))
	if err != nil {
		level = logrus.ErrorLevel
	}
	logrus.SetLevel(level)

	if c.PrettyConsole {
		logrus.SetFormatter(&logrus.TextFormatter{
			FullTimestamp:   true,
			ForceColors:     true,
			DisableColors:   false,
			TimestampFormat: "2006-01-02 15:04:05",
		})
	} else {
		logrus.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: "2006-01-02T15:04:05Z07:00",
			PrettyPrint:     false,
		})
	}
}
