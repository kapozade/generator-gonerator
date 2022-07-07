package models

type ApplicationConfiguration struct {
	Port int64 `env:"PORT" envDefault:"5000"`
}
