package youtube_worker

import (
	"fmt"
	"time"
)

func BeginWorkerCron() {
	ticker := GetTickerDuration()

	for {
		select {
		case <-ticker.C:
			fmt.Println("Fetching of youtube data has started via cron")

			//ctx := &gin.Context{}

		}
	}
}

func GetTickerDuration() *time.Ticker {
	//var value int
	//value =
	return time.NewTicker(time.Duration(3) * time.Second)
}
