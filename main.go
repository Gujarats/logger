package logger

import (
	"fmt"
	"log"
)

// this is for customize the output format time
type logWriter struct{}

// this is for customize the output format time
func (writer logWriter) Write(bytes []byte) (int, error) {
	timeFormatted := getColorTime(getTimeNowFormat()) + string(bytes)
	return fmt.Print(timeFormatted)
}

func Debug(prefix, message string) {

	// setup log
	log.SetFlags(0)
	log.SetOutput(new(logWriter))
	debugMessage := getDebugColor(prefix, message)
	log.Println(debugMessage)

	//set logger file output
	dir := currentDir()
	loggerFileName := getFileName()
	file := createLogFile(dir, loggerFileName)
	defer file.Close()
	file.WriteString(getColor(getTimeNowFormat(), prefix, message) + "\n")
}

func getDebugColor(prefix, message string) string {
	prefix = GetColorFormat(Magenta, Faint, prefix)
	message = GetColorFormat(Yellow, Faint, message)
	return prefix + message
}

func getColor(time, prefix, message string) string {
	time = getColorTime(time)
	debugMessage := getDebugColor(prefix, message)
	return time + debugMessage
}
