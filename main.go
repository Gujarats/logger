package logger

import (
	"fmt"
	"log"
)

// Print Debug message and save that message into a file = logger.log
func Debugf(prefix string, message ...interface{}) {
	messageString := fmt.Sprintln(message...)
	//Create debug message with default style
	debugStyle := NewDebugStyle(prefix, messageString)
	timeStyle := NewTimeStyle()

	// setup log
	log.SetFlags(0)
	log.SetOutput(new(customWriter))
	log.Println(debugStyle)

	saveDebugMessage(timeStyle, debugStyle)
}

// Print debug message to os.StdOut without save it to a file
func Debug(prefix string, message ...interface{}) {
	messageString := fmt.Sprintln(message...)
	//Create debug message with default style
	debugStyle := NewDebugStyle(prefix, messageString)

	// setup log
	log.SetFlags(0)
	log.SetOutput(new(customWriter))
	log.Println(debugStyle)
}
