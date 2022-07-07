package models

type HealthCheck struct {
	Name string `json:"name,omitempty" env:"CONTAINER_NAME" envDefault:"local"`
}
