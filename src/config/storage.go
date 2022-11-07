package config

type ConfigPostgres struct {
	Username, Password, Host, Port, Database string
}

func GetPostgersConfig() ConfigPostgres {
	return ConfigPostgres{
		Username: "postgres",
		Password: "postgrespw",
		Host:     "localhost",
		Port:     "49153",
		Database: "postgres",
	}
}
