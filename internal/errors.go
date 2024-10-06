package internal

import (
	log "github.com/sirupsen/logrus"
)

func CheckIfError(err error) {
	if err != nil {
		log.Error(err)
	}
}

