package utils

import (
	"log"
)

func CheckFatal(err error, v ...interface{}) {
	if err != nil {
		MangoLog.fatal(v, err)
	}
}

func checkFatal(err error, v ...interface{}) {
	if err != nil {
		log.Fatalln(v, err)
	}
}
