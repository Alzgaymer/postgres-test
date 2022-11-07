package config

type ConfigPostgres struct {
	Username, Password, Host, Port, Database string
}

var cfg ConfigPostgres = ConfigPostgres{
	Username: "postgres",
	Password: "postgrespw",
	Host:     "localhost",
	Port:     "49155",
	Database: "postgres",
}

func GetPostgersConfig() *ConfigPostgres {
	return &cfg
}
