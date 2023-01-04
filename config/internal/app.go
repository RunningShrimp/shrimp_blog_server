package internal

type config struct {
	Server   server   `yaml:"server"`
	Database database `yaml:"database"`
	Redis    redis    `yaml:"redis"`
}

type App struct {
	Config config `yaml:"app"`
}
