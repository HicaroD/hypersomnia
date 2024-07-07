package logger

import (
	"log"
	"os"
)

var (
	logFile *os.File
	Info    *log.Logger
	Error   *log.Logger
)

func InitLogFile() error {
	// TODO: get path to configuration directory
	var err error

	path := "log.txt"
	logFile, err = os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	Info = getLogger(logFile, "INFO: ")
	Error = getLogger(logFile, "ERROR: ")
	return nil
}

func Close() error {
	return logFile.Close()
}

func getLogger(file *os.File, prefix string) *log.Logger {
	return log.New(file, prefix, log.Ldate|log.Ltime|log.Lshortfile)
}
