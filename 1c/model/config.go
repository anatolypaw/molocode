package model

// Config структура для хранения конфигурационных данных из YAML.
type Config struct {
	MainServer    Server
	ReserveServer Server
	Storage       Server
	GetPeriod     int // период запроса кодов в секундах
}

// Server структура для конфигурации сервера.
type Server struct {
	Host string `yaml:""`
	Port int    `yaml:""`
}
