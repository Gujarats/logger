package logger

import (
	"os"
	"path/filepath"
)

// carefull with the permision if we got permission denied will loop forever
func createLogFile(filePath, fileName string) *os.File {
	f, err := os.OpenFile(filePath+fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		// error file is not created
		err = os.MkdirAll(filePath, os.ModePerm)
		if err != nil {
			panic(err)
		}
		createLogFile(filePath, fileName)
	}

	return f
}

// return absolute path in the current directory
func currentDir() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	dir := filepath.Dir(ex)

	return dir + "/"
}

// Get filename for file logger name
func getFileName() string {
	return "logger.log"
}
