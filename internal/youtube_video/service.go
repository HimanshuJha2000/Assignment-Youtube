package youtube_video

import (
	"math"
)

type Service struct{}

func (service Service) FetchVideoByPageNumber(pageNo int) (map[string]interface{}, error) {

	videoData, err := GetAllVideos()
	lowBound := (pageNo - 1) * 8
	highBound := math.Min(float64(pageNo*8), float64(len(videoData)))
	videoData = videoData[lowBound:int(highBound)]

	result := map[string]interface{}{}

	result["Videos_count"] = len(videoData)
	result["Videos"] = videoData

	return result, err
}

func (service Service) SearchVideoByQuery(title string, description string) (map[string]interface{}, error) {
	var (
		videoData VideoDataModel
		err       error
	)
	if title != "" {
		videoData, err = SearchByTitle(title)
	} else {
		videoData, err = SearchByDescription(description)
	}

	result := map[string]interface{}{}

	if err != nil || videoData.Title == "" {
		result["Found_Video"] = "Video doesn't exist in the database! Try with different title/description"
	} else {
		result["Found_Video"] = videoData
	}
	return result, err
}
