package youtube_video

import (
	"encoding/json"
	"fmt"
	"github.com/gojektech/heimdall/httpclient"
	"github.com/razorpay/MachineRound/internal/api_key"
	"github.com/razorpay/MachineRound/internal/config"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	youtubeConfig config.YoutubeConfig
)

type YoutubeService struct{}

func (YService *YoutubeService) Initialize() {
	youtubeConfig = config.GetConfig().Youtube
}

func BeginWorkerCron() {
	ticker := time.NewTicker(time.Duration(youtubeConfig.TickerTime) * time.Second)

	var (
		response map[string]interface{}
		err      error
		code     int
	)
	for {
		select {
		case <-ticker.C:
			logrus.Println("Fetching of youtube data has started via cron")

			response, err, code = makeCallToYoutubeApi()
			if code == http.StatusOK {
				SaveInDatabase(response)
			} else {
				logrus.Println("Received error while calling Youtube API! Error is ", err)
				continue
			}
		}
	}
}

func makeCallToYoutubeApi() (map[string]interface{}, error, int) {
	var apiKeyObj api_key.ApiKeyModel
	apiKey := apiKeyObj.GetApiKey()

	if apiKey == "" {
		logrus.Println("No API key found! Hitting Youtube API has been stopped! ")
		return nil, fmt.Errorf("no Valid/Available API key"), http.StatusForbidden
	}

	URL := config.GetYoutubeURLRequestEndpoint(youtubeConfig.Endpoint, apiKey, youtubeConfig.MaxResults, youtubeConfig.Query)
	req, err := http.NewRequest(http.MethodGet, URL, nil)
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		return nil, err, http.StatusBadRequest
	}

	apiKeyObj.CloseIfExceededThreshold()

	client := httpclient.NewClient(httpclient.WithHTTPTimeout(time.Duration(15) * time.Second))
	response, rerr := client.Do(req)

	responseBody, _ := ioutil.ReadAll(response.Body)

	responseMap := make(map[string]interface{})
	_ = json.Unmarshal(responseBody, &responseMap)

	if response != nil {
		defer response.Body.Close()
	}

	if rerr != nil {
		return nil, err, http.StatusInternalServerError
	}
	return responseMap, err, response.StatusCode
}

func SaveInDatabase(response map[string]interface{}) {
	var YResp YoutubeResponse
	respBytes, _ := json.Marshal(response)
	_ = json.Unmarshal(respBytes, &YResp)

	for _, item := range YResp.ItemArray {
		var videoObj VideoDataModel

		videoObj.Title = item.Snippet.Title
		videoObj.Description = item.Snippet.Description
		videoObj.ChannelTitle = item.Snippet.ChannelTitle
		videoObj.ThumbnailUrl = item.Snippet.Thumbnail.DefaultValue.URL

		unixTime, _ := time.Parse(time.RFC3339, item.Snippet.PublishedAt)
		videoObj.PublishedAt = unixTime.Unix()

		err := videoObj.Create()

		if err == nil {
			logrus.Println("New video has been added to the database! Video title : ", videoObj.Title)
		} else {
			logrus.Println("This video is already present in the database!")
		}
	}
}
