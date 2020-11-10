package log

import (
	log "github.com/sirupsen/logrus"
)

// Init use logrus to log.
// this sample can log the request with field and do something.
// then can use log.Info() in this project.
func Init(requestId string) *log.Entry {
	// Log as JSON instead of the default ASCII formatter.
	//log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	//log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	//log.SetLevel(log.WarnLevel)
	return log.WithFields(log.Fields{
		"requestId": requestId,
	})
}
