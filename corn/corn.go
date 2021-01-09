package corn

import (
	"github.com/robfig/cron/v3"
	"smya/mytime"
)

// 定时获取时间
func TimeTask() {
	crontab := cron.New()
	_, _ = crontab.AddFunc("* * * * *", mytime.SetHour) // 添加定时任务
	crontab.Start()
	select {}
}
