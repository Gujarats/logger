package logger

import (
	"fmt"
	"log"
	"time"
)

const timeFormat = "2006-01-02::3:04PM::"

type logWriter struct{}

// this is for customize the output format time
func (writer logWriter) Write(bytes []byte) (int, error) {
	return fmt.Print(time.Now().UTC().Format(timeFormat) + string(bytes))
}

func Debug(prefix, message string) {

	timeNow := getTimeFormat()

	// setup log
	log.SetFlags(0)
	log.SetOutput(new(logWriter))
	log.Println(prefix + message)

	//set logger file output
	dir := currentDir()
	loggerFileName := getFileName()
	file := createLogFile(dir, loggerFileName)
	defer file.Close()
	file.WriteString(timeNow + prefix + message + "\n")
}

func getTimeFormat() string {
	t := time.Now()
	return t.Format(timeFormat)
}
