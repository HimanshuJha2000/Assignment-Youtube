package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/razorpay/MachineRound/internal/api_key"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

type ApiPodController struct {
	apiService api_key.Service
}

var ApiPod ApiPodController

type ApiKeyStruct struct {
	ApiKey string `json:"api_key"`
}

func (controller ApiPodController) CreateAPIKey(ctx *gin.Context) {
	val, _ := ioutil.ReadAll(ctx.Request.Body)

	var obj ApiKeyStruct
	json.Unmarshal(val, &obj)

	err := controller.apiService.AddApiKey(obj.ApiKey)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, "Error : Bad Request! API key already exists!")
	} else {
		logrus.Println("New API key has been added to database")
		ctx.JSON(http.StatusOK, "API key added successfully!")
	}
}
