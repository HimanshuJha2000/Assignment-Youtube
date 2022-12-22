package video

import (
	"encoding/json"
	"fmt"
	"github.com/razorpay/MachineRound/internal/api_key"
	"github.com/razorpay/MachineRound/internal/config"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	youtubeWorkerClient *http.Client
	youtubeConfig       config.YoutubeConfig
)

type YoutubeService struct{}

func (YService *YoutubeService) Initialize() {
	youtubeConfig = config.GetConfig().Youtube
	youtubeWorkerClient = config.GetNewHTTPClient(&youtubeConfig.HttpClient)
}

func BeginWorkerCron() {
	ticker := time.NewTicker(time.Duration(youtubeConfig.TickerTime) * time.Second)

	for {
		select {
		case <-ticker.C:
			fmt.Println("Fetching of youtube data has started via cron")

			//makeCallToYoutubeApi()

			//case <-ctx.Done():
			//	ticker.Stop()
			//	return
		}
	}
}

func makeCallToYoutubeApi() error {
	URL := config.GetYoutubeURLRequestEndpoint(youtubeConfig.Endpoint, api_key.GetApiKey(), youtubeConfig.MaxResults, youtubeConfig.Query)

	req, err := http.NewRequest(http.MethodGet, URL, nil)
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		return fmt.Errorf("error in making new http request, got error %+v", err)
	}

	response, rerr := youtubeWorkerClient.Do(req)

	responseBody, readerr := ioutil.ReadAll(response.Body)

	if readerr != nil {
		fmt.Errorf("")
	}
	responseMap := make(map[string]interface{})
	_ = json.Unmarshal(responseBody, &responseMap)

	if response != nil {
		defer response.Body.Close()
	}

	if err != nil {
		fmt.Errorf("error is ", rerr)
	}
	return nil
}
