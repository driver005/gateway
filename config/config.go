package config

import "github.com/spf13/viper"

type Server struct {
	Port             int    `mapstructure:"port"`
	Host             string `mapstructure:"host"`
	ServerName       string `mapstructure:"server_name"`
	RequestBodyLimit int    `mapstructure:"request_body_limit"`
	Timeout          int    `mapstructure:"timeout"`
	AdminCors        string `mapstructure:"admin_cors"`
	StoreCors        string `mapstructure:"store_cors"`
}

type Database struct {
	Type              string `mapstructure:"type"`
	Username          string `mapstructure:"username"`
	Password          string `mapstructure:"password"`
	Host              string `mapstructure:"host"`
	Port              int    `mapstructure:"port"`
	DBname            string `mapstructure:"dbname"`
	PoolSize          int    `mapstructure:"pool_size"`
	ConnectionTimeout int    `mapstructure:"connection_timeout"`
}

type Secrets struct {
	MasterEncKey string `mapstructure:"master_enc_key"`
	JwtSecret    string `mapstructure:"jwt_secret"`
}

type Applictaion struct {
	Preload  bool     `mapstructure:"preload"`
	Features []string `mapstructure:"features"`
}

type Migration struct {
	Active bool `mapstructure:"active"`
}

type Logger struct {
	Development bool   `mapstructure:"development"`
	Level       string `mapstructure:"level"`
	Encoding    string `mapstructure:"encoding"`
}

type Config struct {
	Server         Server      `mapstructure:"server"`
	MasterDatabase Database    `mapstructure:"master_database"`
	Secrets        Secrets     `mapstructure:"secrets"`
	Applictaion    Applictaion `mapstructure:"applictaion"`
	Logger         Logger      `mapstructure:"logger"`
	Migration      Migration   `mapstructure:"migration"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config *Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("default")
	viper.SetConfigType("toml")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
