package response

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

// LogSuccessfulResponse ...
func LogSuccessfulResponse(r *http.Request, data interface{}) {

	logrus.WithFields(logrus.Fields{
		"host":       r.Host,
		"address":    r.RemoteAddr,
		"method":     r.Method,
		"requestURI": r.RequestURI,
		"proto":      r.Proto,
		"useragent":  r.UserAgent(),
	}).Debug(data)
}
