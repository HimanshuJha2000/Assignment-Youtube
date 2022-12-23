package api_key

type Service struct{}

func (service Service) AddApiKey(api_key string) error {
	apiModelObj := ApiKeyModel{}

	apiModelObj.ApiKey = api_key
	apiModelObj.IsValid = 1

	err := apiModelObj.Create()

	return err
}
