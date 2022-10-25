package internal

import "log"

func PanicOnError(err error, message string) {
	if err != nil {
		log.Panicf("%s: %s", message, err)
	}
}
