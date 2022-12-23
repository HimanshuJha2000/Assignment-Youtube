package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/razorpay/MachineRound/internal/youtube_video"
	"io/ioutil"
	"net/http"
	"strconv"
)

type YoutubeController struct {
	youtubeService youtube_video.Service
}

var YoutubePod YoutubeController

type YoutubeSearchQuery struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (controller YoutubeController) FetchVideoDetails(ctx *gin.Context) {

	pageNo := ctx.Params.ByName("page_no")
	pageNoValue, _ := strconv.Atoi(pageNo)

	result, err := controller.youtubeService.FetchVideoByPageNumber(pageNoValue)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, result)
	} else {
		ctx.JSON(http.StatusOK, result)
	}
}

func (controller YoutubeController) SearchVideo(ctx *gin.Context) {

	validInput, _ := ioutil.ReadAll(ctx.Request.Body)

	var obj YoutubeSearchQuery
	json.Unmarshal(validInput, &obj)

	result, err := controller.youtubeService.SearchVideoByQuery(obj.Title, obj.Description)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, result)
	} else {
		ctx.JSON(http.StatusOK, result)
	}
}
