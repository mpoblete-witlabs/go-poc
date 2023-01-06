package utils

type Config struct {
	Port        string
	MysqlConfig MysqlConfig
}

type MysqlConfig struct {
	Password string
	Db       string
	User     string
	Port     string
	Host     string
}

func GetConfig() Config {

	config := Config{
		Port: "",
		MysqlConfig: MysqlConfig{
			Password: "admin",
			Db:       "poc-go",
			User:     "root",
			Port:     "3306",
			Host:     "localhost",
		},
	}
	return config
}
