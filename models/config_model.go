package models

var RuntimeConf = RuntimeConfig{}

type RuntimeConfig struct {
	GolangProfile   string     `yaml:"golangProfile"`
	DbMigrationPath string     `yaml:"dbMigrationPath"`
	Datasource      Datasource `yaml:"datasource"`
	Server          Server     `yaml:"server"`
}

type Datasource struct {
	DbType   string `yaml:"dbType"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	UserName string `yaml:"userName"`
	Password string `yaml:"password"`
	DbName   string `yaml:"dbName"`
}

type Server struct {
	Port string `yaml:"port"`
}
