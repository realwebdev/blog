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
	DSN string
}
