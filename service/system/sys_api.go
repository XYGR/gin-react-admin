package system

import (
	"errors"
	"gin-react-admin/global"
	"gin-react-admin/model/system"
	"gorm.io/gorm"
)

type ApiService struct{}

var ApiServiceApp = new(ApiService)

//@function: CreateApi
//@description: 创建Api接口
//@param api system.SysApi
//@return error

func (apiService *ApiService) CreateApi(api system.SysApi) (err error) {
	if !errors.Is(global.GRA_DB.Where("path = ? AND method = ?", api.Path, api.Method).First(&system.SysApi{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同api")
	}
	return global.GRA_DB.Create(&api).Error
}
