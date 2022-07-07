package services

import (
	"<%=appname%>/models"

	"github.com/caarlos0/env/v6"
)

type ConfigurationServiceInterface interface {
	LoadHealthCheckConfig() (*models.HealthCheck, *models.ApplicationErrors)
	GetApplicationConfiguration() (*models.ApplicationConfiguration, *models.ApplicationErrors)
}

type configurationService struct {
}

var (
	ConfigurationService ConfigurationServiceInterface
)

func init() {
	ConfigurationService = &configurationService{}
}

func (c *configurationService) LoadHealthCheckConfig() (*models.HealthCheck, *models.ApplicationErrors) {
	var config models.HealthCheck
	if err := env.Parse(&config); err != nil {
		var applicationErrors []models.ApplicationError
		applicationError := models.ApplicationError{
			Message: err.Error(),
		}
		applicationErrors = append(applicationErrors, applicationError)

		return nil, &models.ApplicationErrors{
			Errors: applicationErrors,
		}
	}

	return &config, nil
}

func (c *configurationService) GetApplicationConfiguration() (*models.ApplicationConfiguration, *models.ApplicationErrors) {
	var config models.ApplicationConfiguration
	if err := env.Parse(&config); err != nil {
		var applicationErrors []models.ApplicationError
		applicationError := models.ApplicationError{
			Message: err.Error(),
		}
		applicationErrors = append(applicationErrors, applicationError)

		return nil, &models.ApplicationErrors{
			Errors: applicationErrors,
		}
	}

	return &config, nil
}
