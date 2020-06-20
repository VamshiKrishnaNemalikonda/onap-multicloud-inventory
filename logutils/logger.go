package logutils

import (
        log "github.com/sirupsen/logrus"
)

//Fields is type that will be used by the calling function
type Fields map[string]interface{}

func init() {
        // Log as JSON instead of the default ASCII formatter.
        log.SetFormatter(&log.JSONFormatter{})
}

// Error uses the fields provided and logs
func Error(msg string) {
        log.Error(msg)
}

func Info(msg string) {
        log.Info(msg)
}

func Debug(msg string) {
        log.Debug(msg)
}


// Warn uses the fields provided and logs
func Warn(msg string, fields Fields) {
        log.WithFields(log.Fields(fields)).Warn(msg)
}



