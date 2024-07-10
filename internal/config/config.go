package config

type Config struct {
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`
	Database struct {
		Driver      string `yaml:"driver"`
		Address     string `yaml:"address"`
		Username    string `yaml:"username"`
		Password    string `yaml:"password"`
		Port        string `yaml:"port"`
		Database    string `yaml:"database"`
		IsPopulated bool   `yaml:"isPopulated"`
		IsMigrate   bool   `yaml:"isMigrate"`
	} `yaml:"database"`
	Cache struct {
		Drive    string `yaml:"drive"`
		Address  string `yaml:"address"`
		Port     string `yaml:"port"`
		DbNum    string `yaml:"dbNum"`
		Password string `yaml:"password"`
	} `yaml:"cache"`
}
