package utils

import (
	"log"
)

// -> CheckErr: check if an error occurred.
func CheckErr(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %v", msg, err)
	}
}

// -> CheckErrAbortProgram: if an error occurred, panic is called.
func CheckErrAbortProgram(err error, msg string) {
	if err != nil {
		panic(msg + ": " + err.Error())
	}
}
