package api_key

import (
	"fmt"
	"github.com/razorpay/MachineRound/internal/providers/database"
)

type ApiKeyModel struct {
	ID      string `json:"id"`
	ApiKey  string `json:"api_key"`
	IsValid bool   `json:"is_valid"`
}

func (Am *ApiKeyModel) Create() error {
	err := database.Client().Create(&Am).Error
	return err
}

func (Am *ApiKeyModel) Update(apiKey string) error {
	err := database.Client().Model(Am).Where("api_key = ?", apiKey).Updates(map[string]interface{}{"is_valid": Am.IsValid}).Error
	return err
}

func GetApiKey() string {
	var AmObj ApiKeyModel
	err := database.Client().Where("is_valid = ?", true).First(&AmObj).Error
	if err != nil {
		fmt.Errorf(err.Error())
		return "Error occurred while fetching API key from database"
	}
	return AmObj.ApiKey
}
