package system

import (
	"gin-react-admin/global"
	"gin-react-admin/model/system"
)

type OperationRecordService struct{}

//@function: CreateSysOperationRecord
//@description: 创建操作记录
//@param: sysOperationRecord system.SysOperationRecord
//@return: err error

func (operationRecordService OperationRecordService) CreateSysOperationRecord(sysOperationRecord system.SysOperationRecord) (err error) {
	err = global.GRA_DB.Create(&sysOperationRecord).Error
	return err
}
