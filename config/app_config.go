package config

var AppConf = AppConfig{}

type AppConfig struct {
	DbMigrationPath string       `mapstructure:"dbMigrationPath"`
	Datasource      Datasource   `mapstructure:"datasource"`
	ServerConfig    ServerConfig `mapstructure:"serverConfig"`
	ZapConfig       LogConfig    `mapstructure:"zapConfig"`
	LorusConfig     LogConfig    `mapstructure:"logrusConfig"`
	Log             LogConfig    `mapstructure:"logConfig"`
}

type Datasource struct {
	DbType string `mapstructure:"dbType"`
	Dsn    string `mapstructure:"dsn"`
}

type ServerConfig struct {
	HttpServerAddress string `mapstructure:"httpServerAddress"`
}

// LogConfig represents logger handler
// Logger has many parameters can be set or changed. Currently, only three are listed here. Can add more into it to
// fits your needs.
type LogConfig struct {
	// log library name
	Code string `mapstructure:"code"`
	// log level
	Level string `mapstructure:"level"`
	// show caller in log message
	EnableCaller bool `mapstructure:"enableCaller"`
}
