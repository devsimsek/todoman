package utils

import "log"

// Error function to log an error and exit the program
func Error(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// ErrorInfo function to log an error message and exit the program
func ErrorInfo(err error, message string, args ...interface{}) {
	if err != nil {
		log.Fatalf(message, args...)
	}
}
