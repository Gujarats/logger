package logger

import (
	"log"
)

func Debug(prefix, message string) {
	//Create debug message with default style
	debugStyle := NewDebugStyle(prefix, message)
	timeStyle := NewTimeStyle()

	// setup log
	log.SetFlags(0)
	log.SetOutput(new(customWriter))
	log.Println(debugStyle)

	saveDebugMessage(timeStyle, debugStyle)
}
