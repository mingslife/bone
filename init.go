package bone

import (
	"log"
	"os"
)

func init() {
	// unnecessary to show file name in package infra
	// log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ldate | log.Ltime)
}
