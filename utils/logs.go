package utils

import (
	"log"
	"os"
)

type LogsError interface {
	InsertLogsError(LogsError error)
}

type logsError struct{}

func NewLogsError() LogsError {
	return &logsError{}
}

func (l *logsError) InsertLogsError(logError error) {

	f, err := os.OpenFile("logs/errors.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	log.SetOutput(f)
	log.Printf("%v", logError)

}
