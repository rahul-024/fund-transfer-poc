package config

var AppConf = AppConfig{}

type AppConfig struct {
	DbMigrationPath string       `yaml:"dbMigrationPath"`
	Datasource      Datasource   `yaml:"datasource"`
	ServerConfig    ServerConfig `yaml:"serverConfig"`
}

type Datasource struct {
	DbType string `yaml:"dbType"`
	Dsn    string `yaml:"dsn"`
}

type ServerConfig struct {
	HttpServerAddress string `yaml:"httpServerAddress"`
}
