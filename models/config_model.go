package models

var RuntimeConf = RuntimeConfig{}

type RuntimeConfig struct {
	DbMigrationPath string     `yaml:"dbMigrationPath"`
	Datasource      Datasource `yaml:"datasource"`
	Server          Server     `yaml:"server"`
}

type Datasource struct {
	DbType string `yaml:"dbType"`
	Dsn    string `yaml:"dsn"`
}

type Server struct {
	HttpServerAddress string `yaml:"httpServerAddress"`
}
