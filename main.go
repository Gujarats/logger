package logger

import (
	"log"
)

// Print Debug message and save that message into a file = logger.log
func Debugf(prefix, message string) {
	//Create debug message with default style
	debugStyle := NewDebugStyle(prefix, message)
	timeStyle := NewTimeStyle()

	// setup log
	log.SetFlags(0)
	log.SetOutput(new(customWriter))
	log.Println(debugStyle)

	saveDebugMessage(timeStyle, debugStyle)
}

// Print debug message to os.StdOut without save it to a file
func Debug(prefix, message string) {
	//Create debug message with default style
	debugStyle := NewDebugStyle(prefix, message)

	// setup log
	log.SetFlags(0)
	log.SetOutput(new(customWriter))
	log.Println(debugStyle)
}
