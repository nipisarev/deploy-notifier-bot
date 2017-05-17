package middlewares

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"fmt"
	"time"
	"io"
	"io/ioutil"
	"net/http"
)

var req []byte

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func Logger() gin.HandlerFunc {
	return log(gin.DefaultWriter)
}

func log(out io.Writer, notlogged ...string) gin.HandlerFunc {

	var skip map[string]struct{}

	if length := len(notlogged); length > 0 {
		skip = make(map[string]struct{}, length)

		for _, path := range notlogged {
			skip[path] = struct{}{}
		}
	}

	return func(c *gin.Context) {
		start := time.Now()
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw
		path := c.Request.URL.Path
		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		if method == "POST" {
			req, _ = GetBody(c.Request)
		}

		c.Next()
		end := time.Now()
		latency := end.Sub(start)

		fmt.Fprintf(out, "%v [Request] %s |status: %3d | %s | %s | body: %s\n"+
			"%v [Response] %13v | body: %s\n",
			end.Format("2006/01/02 - 15:04:05"),
			method,
			statusCode,
			clientIP,
			path,
			req,
			end.Format("2006/01/02 - 15:04:05"),
			latency,
			blw.body.String(),
		)
	}
}

func GetBody(req *http.Request) ([]byte, error) {
	reqBody, err := ioutil.ReadAll(io.LimitReader(req.Body, 1048576))
	req.Body.Close()
	req.Body = ioutil.NopCloser(bytes.NewReader(reqBody))

	return reqBody, err
}
