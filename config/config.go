package config

type Config struct {
	App      App
	Server   Server
	Database Database
}

type App struct {
	Name string
}

type Server struct {
	Host string
	Port string
}

type Database struct {
	UserName string
	Password string
	Host     string
	Port     string
	DBName   string
}
