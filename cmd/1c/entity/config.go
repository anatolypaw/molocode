package entity

// Server структура для конфигурации сервера.
type Server struct {
	IP   string `yaml:""`
	Port int    `yaml:""`
}

// Config структура для хранения конфигурационных данных из YAML.
type Config struct {
	Main_exchangemarks   Server
	Reserve_exhangemarks Server
	Hub                  Server
	GetPeriod            int // период запроса кодов в секундах
}
