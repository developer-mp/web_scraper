package internal

import (
	"bytes"

	"github.com/gin-gonic/gin"
	"github.com/sebest/logrusly"
	"github.com/sirupsen/logrus"
)

type ResponseLogger struct {
    gin.ResponseWriter
    Body *bytes.Buffer
}

func (r *ResponseLogger) Write(data []byte) (int, error) {
    r.Body.Write(data)
    return r.ResponseWriter.Write(data)
}

func HandleAPIEndpoint(c *gin.Context, l *logrus.Logger, h *logrusly.LogglyHook, host string, handlerFunc func(c *gin.Context)) {
    responseBody := &bytes.Buffer{}
    responseLogger := &ResponseLogger{c.Writer, responseBody}
    c.Writer = responseLogger

    handlerFunc(c)

    l.WithFields(logrus.Fields{"location": host, "response": responseBody.String()}).Warn("Response from backend server")
    h.Flush()
}
