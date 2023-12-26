package model

// Config структура для хранения конфигурационных данных из YAML.
type Config struct {
	MainServer    Server
	ReserveServer Server
}

// Server структура для конфигурации сервера.
type Server struct {
	Host string `yaml:""`
	Port int    `yaml:""`
}
