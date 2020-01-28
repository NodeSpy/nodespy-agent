package logging

import (
	"io"
	"os"

	"github.com/rs/zerolog"
)

type Config struct {
	Debug   bool
	LogFile string
}

type Logger struct {
	*zerolog.Logger
}

func Configure(config Config) *Logger {
	var writers []io.Writer

	writers = append(writers, zerolog.ConsoleWriter{Out: os.Stderr}) // Console out

	if config.LogFile != "" {
		writers = append(writers, fileLogger(config.LogFile))
	}

	mw := io.MultiWriter(writers...)

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if config.Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
	logger := zerolog.New(mw).With().Timestamp().Logger()

	return &Logger{
		Logger: &logger,
	}

}

func fileLogger(fileName string) io.Writer {
	logFile, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}

	return logFile
}
