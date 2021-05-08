package log

import (
	"os"

	"github.com/sirupsen/logrus"
)

const (
	LogLevelEnvVar = "LOG_LEVEL"
	LogJSONEnvVar  = "LOG_JSON"
)

// New is a project global creator of logger
func New(name string) *Logger {
	l := logrus.New()
	l.SetOutput(os.Stdout)
	l.SetReportCaller(true)

	if os.Getenv(LogJSONEnvVar) != "" {
		l.SetFormatter(&logrus.JSONFormatter{})
	} else {
		l.SetFormatter(&logrus.TextFormatter{
			DisableTimestamp: true,
			ForceColors:      true,
		})
	}

	lev, ok := os.LookupEnv(LogLevelEnvVar)
	if !ok {
		l.Infof("%s not set, using default: %s", LogLevelEnvVar, logrus.InfoLevel)
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
		name: name,
		logger: l.WithFields(logrus.Fields{
			"app": name,
		}),
	}
}

// Logger is the internal logrus abstraction
type Logger struct {
	name   string
	logger *logrus.Entry
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