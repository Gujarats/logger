package logger

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
)

const calldepth = 2

// setup logger
var myLogger *log.Logger

func init() {
	myLogger = log.New(os.Stderr, "", log.Lshortfile)
}

// Print Debug message and save that message into a file = logger.log
// the file location logger.lo would be in the root of the project directory
func Debugf(prefix string, message ...interface{}) {
	messageString := fmt.Sprintln(message...)

	//Create debug message with default style
	debugStyle := NewDebugStyle(prefix, messageString)
	timeStyle := NewTimeStyle()

	// output logger
	myLogger.SetOutput(new(customWriter))
	myLogger.Output(calldepth, debugStyle)

	// save output logger
	lineError := getLineError()
	saveDebugMessage(timeStyle, lineError+debugStyle)
}

func getLineError() string {
	var fileShortend string
	_, file, line, ok := runtime.Caller(calldepth)
	if !ok {
		return ""
	}

	for i := len(file) - 1; i > 0; i-- {
		if file[i] == '/' {
			fileShortend = file[i+1:]
			break
		}
	}

	lineStr := strconv.Itoa(line)

	return fileShortend + ":" + lineStr + ": "
}

// Print debug message to os.StdOut without save it to a file
func Debug(prefix string, message ...interface{}) {
	messageString := fmt.Sprintln(message...)

	//Create debug message with default style
	debugStyle := NewDebugStyle(prefix, messageString)

	// output logger
	myLogger.SetOutput(new(customWriter))
	myLogger.Output(calldepth, debugStyle)
}
