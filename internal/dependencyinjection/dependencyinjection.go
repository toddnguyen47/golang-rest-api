package dependencyinjection

import (
	"golang_rest_api/internal/restapi"

	"github.com/spf13/viper"
)

type DependencyInjection struct {
	RestApi restapi.RestApi
}

func NewDependencyInjection() DependencyInjection {
	restApi := restapi.NewRestApi(restapi.Config{
		Url: viper.GetString("URL_PATH"),
	})

	return DependencyInjection{
		RestApi: restApi,
	}
}
