package server

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/gin-gonic/gin"
)

func StartGinServer(port int) {
	r := gin.Default()
	r.Any("/:anypath", func(c *gin.Context) {
		reqByte, _ := httputil.DumpRequest(c.Request, true)
		log.Printf("req: %s", reqByte)
		bodyByte, _ := ioutil.ReadAll(c.Request.Body)
		for k := range c.Request.Header {
			c.Header(k, c.Request.Header.Get(k))
		}
		c.Data(http.StatusOK, c.Request.Header.Get("Content-Type"), bodyByte)
	})
	r.Run(fmt.Sprintf(":%d", port))
}
