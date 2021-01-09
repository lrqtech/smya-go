package corn

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"smya/mytime"
)

// 定时获取时间
func TimeTask() {
	crontab := cron.New()
	_, err := crontab.AddFunc("* * * * *", mytime.SetHour) // 添加定时任务
	if err != nil {
		fmt.Println(err)
	}
	crontab.Start()
	select {}
}
