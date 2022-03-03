package middleware

import (
	"bytes"
	"encoding/json"
	"gin-react-admin/global"
	"gin-react-admin/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var operationRecordService = service.ServiceGroupApp.SystemServiceGroup.OperationRecordService

func OperationRecord() gin.HandlerFunc {
	return func(c *gin.Context) {
		var body []byte
		var userId int
		if c.Request.Method != http.MethodGet {
			var err error
			body, err = ioutil.ReadAll(c.Request.Body)
			if err != nil {
				global.GRA_LOG.Error("read body from request error:", zap.Error(err))
			} else {
				c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
			}
		} else {
			query := c.Request.URL.RawQuery
			query, _ = url.QueryUnescape(query)
			split := strings.Split(query,"&")
			m := make(map[string]string)
			for _, v := range split {
				kv := strings.Split(v,"=")
				if len(kv) == 2 {
					m[kv[0]] = m[kv[1]]
				}
			}
			body, _ = json.Marshal(&m)
		}
		claims, _ := utils.
	}
}
