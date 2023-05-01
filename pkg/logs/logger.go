package logs

import (
	log "github.com/sirupsen/logrus"
	"os"
)

var entry *log.Entry

type Logger struct {
	*log.Entry
}

func GetLogger() *Logger {
	return &Logger{
		entry,
	}
}

func init() {
	logger := log.New()
	logger.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
	logger.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
	entry = log.NewEntry(logger)
}
