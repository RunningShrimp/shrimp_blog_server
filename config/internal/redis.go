package internal

type redis struct {
	Database int    `yaml:"database"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
}
