package config

import dbRepo "go_template_project/internal/repository"

type (
	EnvVars struct {
		ServerHost       string `envconfig:"server_host"`
		ServerPort       int    `envconfig:"server_port"`
		ServerAllowCors  bool   `envconfig:"server_allow_cors"`
		ServerDebugMode  bool   `envconfig:"server_debug_mode"`
		SwaggerDocs      bool   `envconfig:"swagger_docs"`
		DatabaseHost     string `envconfig:"db_host"`
		DatabasePort     int    `envconfig:"db_port"`
		DatabaseName     string `envconfig:"db_name"`
		DatabaseUsername string `envconfig:"db_username"`
		DatabasePassword string `envconfig:"db_password"`
	}

	serverConfig struct {
		Host        string
		Port        int
		AllowCors   bool
		DebugMode   bool
		SwaggerDocs bool
	}

	Config struct {
		Server     serverConfig
		Repository dbRepo.Config
	}
)

func NewConfig(f EnvVars) Config {
	return Config{
		Server: serverConfig{
			Host:        f.ServerHost,
			Port:        f.ServerPort,
			AllowCors:   f.ServerAllowCors,
			DebugMode:   f.ServerDebugMode,
			SwaggerDocs: f.SwaggerDocs,
		},
		Repository: dbRepo.Config{
			Host:     f.DatabaseHost,
			Port:     f.DatabasePort,
			Name:     f.DatabaseName,
			Username: f.DatabaseUsername,
			Password: f.DatabasePassword,
		},
	}
}
