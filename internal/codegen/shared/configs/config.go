package configs

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	viperConfig = viper.New()
	config      = &Config{}
)

type Config struct {
	ProjectName  string            `yaml:"projectName"`
	TemplatePath string            `yaml:"templatePath"`
	OutputPath   string            `yaml:"outputPath"`
	UsingGit     bool              `yaml:"usingGit"`
	MappingValue map[string]string `yaml:"mappingValue"`
}

func ReadConfig() {
	// check for viper file in current directory
	// if not found, return error
	// if found, read viper file using viper
	viperConfig.SetConfigName("config")
	viperConfig.AddConfigPath("./configs")
	// viper file should be yaml
	viperConfig.SetConfigType("yaml")

	// read into config struct
	err := viperConfig.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viperConfig.Unmarshal(config)
	if err != nil {
		panic(err)
	}
}

func GetConfig() *Config {
	return config
}

func SetFlag(key string, flag *pflag.Flag) {
	viperConfig.BindPFlag(key, flag)
}

func GetFlag(key string) any {
	return viperConfig.Get(key)
}
