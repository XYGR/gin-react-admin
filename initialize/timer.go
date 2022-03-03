package initialize

import (
	"fmt"
	"gin-react-admin/config"
	"gin-react-admin/global"
	"gin-react-admin/utils"
)

func Timer() {
	if global.GRA_CONFIG.Timer.Start {
		for i := range global.GRA_CONFIG.Timer.Detail {
			go func(detail config.Detail) {
				global.GRA_Timer.AddTaskByFunc("clearDB", global.GRA_CONFIG.Timer.Spec, func() {
					err := utils.ClearTable(global.GRA_DB, detail.TableName, detail.CompareField, detail.Interval)
					if err != nil {
						fmt.Println("timer error:", err)
					}
				})
			}(global.GRA_CONFIG.Timer.Detail[i])
		}
	}
}
