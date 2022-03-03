package system

import "gin-react-admin/global"

type JwtBlacklist struct {
	global.GRA_MODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}
