package api_key

import (
	"fmt"
	"github.com/razorpay/MachineRound/internal/constants"
	"github.com/razorpay/MachineRound/internal/providers/database"
)

type ApiKeyModel struct {
	ApiKey     string `json:"api_key"`
	IsValid    int64  `json:"is_valid"`
	UsageCount int64  `json:"usage_count"`
}

func (Am *ApiKeyModel) Create() error {
	err := database.Client().Create(&Am).Error
	return err
}

func (Am *ApiKeyModel) GetApiKey() string {
	err := database.Client().Model(&Am).Where("is_valid = ?", 1).First(&Am).Error
	err = database.Client().Model(&Am).Where("api_key = ?", Am.ApiKey).Updates(map[string]interface{}{"usage_count": Am.UsageCount + 1}).Error
	if err != nil {
		fmt.Errorf(err.Error())
		return Am.ApiKey
	}
	return Am.ApiKey
}

func (Am *ApiKeyModel) CloseIfExceededThreshold() {
	if Am.UsageCount > constants.API_Key_Threshold {
		err := database.Client().Model(&Am).Where("api_key = ?", Am.ApiKey).Updates(map[string]interface{}{"is_valid": 0}).Error
		fmt.Errorf(err.Error())
	}
}
