package video

import (
	"github.com/razorpay/MachineRound/internal/providers/database"
)

type YoutubeResponse struct {
	ItemArray []Items `json:"items"`
}

type Items struct {
	Kind    string `json:"kind"`
	Etag    string `json:"etag"`
	Snippet Video  `json:"snippet"`
}

type Video struct {
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	PublishingTime string    `json:"publishedAt"`
	ChannelTitle   string    `json:"channelTitle"`
	Thumbnail      Thumbnail `json:"thumbnail"`
}

type Thumbnail struct {
	DefaultValue Default `json:"default"`
}

type Default struct {
	URL    string `json:"url"`
	Width  int64  `json:"width"`
	Height int64  `json:"height"`
}

type VideoDataModel struct {
	ID             string `json:"id"`
	Title          string `json:"title"`
	Description    string `json:"description"`
	PublishingTime int64  `json:"published_at"`
	ChannelTitle   string `json:"channel_title"`
	ThumbnailUrl   string `json:"thumbnail_url"`
}

func (dm *VideoDataModel) Create() error {
	err := database.Client().Create(&dm).Error
	return err
}

func (dm *VideoDataModel) Get() error {
	return nil
}
