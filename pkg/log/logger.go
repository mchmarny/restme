package log

import (
	"os"

	"github.com/sirupsen/logrus"
)

const (
	nameDefault    = "test"
	levelDefault   = "debug"
	versionDefault = "v0.0.1v"
)

func Default() *Logger {
	return New(nameDefault, versionDefault, levelDefault, false)
}

// New is a project global creator of logger
func New(name, version, level string, json bool) *Logger {
	l := logrus.New()
	l.SetOutput(os.Stdout)
	l.SetReportCaller(true)

	if json {
		l.SetFormatter(&logrus.JSONFormatter{})
	} else {
		l.SetFormatter(&logrus.TextFormatter{
			DisableTimestamp: true,
			ForceColors:      true,
		})
	}

	lev, ok := os.LookupEnv(level)
	if !ok {
		l.Infof("%s not set, using default: %s", level, logrus.InfoLevel)
		lev = logrus.InfoLevel.String()
	}

	v, err := logrus.ParseLevel(lev)
	if err != nil {
		l.Errorf("invalid debug level format: %s", lev)
		l.SetLevel(logrus.InfoLevel)
	}
	l.SetLevel(v)
	if !l.IsLevelEnabled(v) {
		l.Errorf("log level not enabled, expected: %s, got: %s", lev, l.GetLevel())
	}

	return &Logger{
		name:    name,
		version: version,
		logger: l.WithFields(logrus.Fields{
			"app":     name,
			"version": version,
		}),
	}
}

// Logger is the internal logrus abstraction
type Logger struct {
	version string
	name    string
	logger  *logrus.Entry
}

// GetAppVersion returns app version.
func (l *Logger) GetAppVersion() string {
	return l.version
}

// GetAppName returns app name.
func (l *Logger) GetAppName() string {
	return l.name
}

// GetLevel returns the logger configured level.
func (l *Logger) GetLevel() logrus.Level {
	return l.logger.Logger.GetLevel()
}

// Info logs a message at level Info.
func (l *Logger) Info(args ...interface{}) {
	l.logger.Log(logrus.InfoLevel, args...)
}

// Infof logs a message at level Info.
func (l *Logger) Infof(format string, args ...interface{}) {
	l.logger.Logf(logrus.InfoLevel, format, args...)
}

// Debug logs a message at level Debug.
func (l *Logger) Debug(args ...interface{}) {
	l.logger.Log(logrus.DebugLevel, args...)
}

// Debugf logs a message at level Debug.
func (l *Logger) Debugf(format string, args ...interface{}) {
	l.logger.Logf(logrus.DebugLevel, format, args...)
}

// Warn logs a message at level Warn.
func (l *Logger) Warn(args ...interface{}) {
	l.logger.Log(logrus.WarnLevel, args...)
}

// Warnf logs a message at level Warn.
func (l *Logger) Warnf(format string, args ...interface{}) {
	l.logger.Logf(logrus.WarnLevel, format, args...)
}

// Error logs a message at level Error.
func (l *Logger) Error(args ...interface{}) {
	l.logger.Log(logrus.ErrorLevel, args...)
}

// Errorf logs a message at level Error.
func (l *Logger) Errorf(format string, args ...interface{}) {
	l.logger.Logf(logrus.ErrorLevel, format, args...)
}

// Fatal logs a message at level Fatal then the process will exit with status set to 1.
func (l *Logger) Fatal(args ...interface{}) {
	l.logger.Fatal(args...)
}

// Fatalf logs a message at level Fatal then the process will exit with status set to 1.
func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.logger.Fatalf(format, args...)
}
