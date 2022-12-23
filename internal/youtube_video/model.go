package youtube_video

import (
	"github.com/razorpay/MachineRound/internal/providers/database"
)

type YoutubeResponse struct {
	ItemArray []Items `json:"items"`
}

type Items struct {
	Snippet Video `json:"snippet"`
}

type Video struct {
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	PublishedAt  string    `json:"publishedAt"`
	ChannelTitle string    `json:"channelTitle"`
	Thumbnail    Thumbnail `json:"thumbnails"`
}

type Thumbnail struct {
	DefaultValue Default `json:"default"`
}

type Default struct {
	URL string `json:"url"`
}

type VideoDataModel struct {
	Title        string `json:"title"`
	Description  string `json:"description"`
	PublishedAt  int64  `json:"published_at"`
	ChannelTitle string `json:"channel_title"`
	ThumbnailUrl string `json:"thumbnail_url"`
}

func (dm *VideoDataModel) Create() error {
	err := database.Client().Model(&dm).Create(&dm).Error
	return err
}

func GetAllVideos() ([]VideoDataModel, error) {
	var VideoItems []VideoDataModel
	err := database.Client().Order("published_at DESC").Find(&VideoItems).Error
	return VideoItems, err
}

func SearchByTitle(title string) (VideoDataModel, error) {
	var VideoItem VideoDataModel
	err := database.Client().Where("title = ?", title).First(&VideoItem).Error
	return VideoItem, err
}

func SearchByDescription(description string) (VideoDataModel, error) {
	var VideoItem VideoDataModel
	err := database.Client().Where("description = ?", description).First(&VideoItem).Error
	return VideoItem, err
}
