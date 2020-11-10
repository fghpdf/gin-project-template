package router

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"time"
)

// Logging a middleware for Gin.
// You can custom logger type for HTTP request in Gin.
func Logging() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		bodyCopy := new(bytes.Buffer)
		// Read the whole body
		_, err := io.Copy(bodyCopy, c.Request.Body)
		if err != nil {
			return // 退出中间件
		}
		bodyData := bodyCopy.Bytes()
		// Replace the body with a reader that reads from the buffer
		c.Request.Body = ioutil.NopCloser(bytes.NewReader(bodyData))

		var requestParams string
		requestParams = string(bodyData)

		c.Next()

		end := time.Now()
		latency := end.Sub(start)
		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		logFields := log.Fields{
			"latency":        fmt.Sprintf("%v", latency),
			"client_ip":      clientIP,
			"host":           c.Request.Host,
			"method":         method,
			"request_id":     c.Writer.Header().Get("X-Request-Id"),
			"request_uri":    c.Request.RequestURI,
			"request_params": requestParams,
			"status_code":    statusCode,
		}

		if c.Errors != nil {
			for _, err := range c.Errors {
				log.Error(err)
			}
		} else {
			log.WithFields(logFields).Info()
		}
	}
}
