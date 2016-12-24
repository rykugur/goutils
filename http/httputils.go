package httputils

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
)

// LogRequest logs the given request to the default logger.
func LogRequest(req *http.Request) {
	log.WithFields(log.Fields{"header": req.Header, "body": req.Body}).Info("")
}
