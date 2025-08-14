package plog

import (
	"log"

	"github.com/Alex-Wolf-7/Kisa/constants"
)

func Periodicf(message string, v ...interface{}) {
	if constants.LOG_PERIODIC_ON {
		log.Printf(message, v...)
	}
}

func Debugf(message string, v ...interface{}) {
	if constants.LOG_DEBUG_ON {
		log.Printf(message, v...)
	}
}

func Infof(message string, v ...interface{}) {
	if constants.LOG_INFO_ON {
		log.Printf(message, v...)
	}
}

func ErrorfWithBackup(backup string, message string, v ...interface{}) {
	if constants.LOG_ERROR_ON {
		log.Printf(message, v...)
	} else {
		log.Print(backup)
	}
}

func FatalfWithCode(code string, message string, v ...interface{}) {
	if constants.LOG_FATAL_ON {
		log.Fatalf(message, v...)
	} else {
		log.Fatalf("Fatal error, closing program. Code: %s\n", code)
	}
}
