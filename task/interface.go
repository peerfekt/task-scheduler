package task

import (
	"log"
)

const DEFAULT_INTERVAL = 500

var DEFAULT_FUNCTION = func() {
	log.Println("Periodic task was executed")
}
