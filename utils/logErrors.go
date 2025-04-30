package utils

import "log"

func LogErrors(err error) {
	log.Fatal(err)
}
