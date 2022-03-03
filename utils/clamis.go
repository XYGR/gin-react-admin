package utils

import (
	systemReq "gin-react-admin/model/system/request"
	"github.com/gin-gonic/gin"
)

func GetClaims(c *gin.Context) (*systemReq.CustomClaims, error) {
	token := c.Request.Header.Get("x-token")
	j := NewJWT()
	claims, err := j.ParseToken(token)
}
