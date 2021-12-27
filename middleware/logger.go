package middleware

import (
	"log"
	"net"
	"time"

	logrustash "github.com/bshuster-repo/logrus-logstash-hook"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Logger(address string) gin.HandlerFunc {

	// Connect to centralized logger.
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatal("Cannot access centralized logger.")
	}

	logger := logrus.New()
	hook := logrustash.New(conn, logrustash.DefaultFormatter(logrus.Fields{"type": "user_service"}))
	logger.Hooks.Add(hook)

	return func(c *gin.Context) {

		// Execution time
		startTime := time.Now()
		c.Next()
		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)

		// Other info.
		reqMethod := c.Request.Method
		reqUri := c.Request.RequestURI
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()

		// Format log.
		logger.WithFields(logrus.Fields{
			"status_code":  statusCode,
			"latency_time": latencyTime,
			"client_ip":    clientIP,
			"req_method":   reqMethod,
			"req_uri":      reqUri,
		}).Info()
	}
}
