package config

type Log struct {
	Level     string `yaml:"level"`
	Output    string `yaml:"output"`
	Name      string `yaml:"name"`
	LogFormat string `yaml:"logFormat"`
}
