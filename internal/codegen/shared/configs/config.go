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
	ProjectName      string            `yaml:"projectName"`
	TemplatePath     string            `yaml:"templatePath"`
	LibraryPath      string            `yaml:"libraryPath"`
	FilePath         string            `yaml:"filePath"`
	OutputPath       string            `yaml:"outputPath"`
	UsingGit         bool              `default:"true" yaml:"usingGit"`
	MappingValue     map[string]string `yaml:"mappingValue"`
	ProjectStructure []FileInfo        `yaml:"projectStructure"`
}

type FileInfo struct {
	Name        string     `yaml:"name"`
	Type        string     `default:"folder" yaml:"type"`
	Children    []FileInfo `yaml:"children"`
	Template    string     `default:"" yaml:"template"`
	File        string     `default:"" yaml:"file"`
	LibraryName string     `default:"" yaml:"libraryName"`
	IsTemplate  bool       `default:"true" yaml:"isTemplate"`
	IsLibrary   bool       `default:"false" yaml:"isLibrary"`
}

func ReadConfig(path string) {
	// check for viper file in current directory
	// if not found, return error
	// if found, read viper file using viper
	viperConfig.SetConfigName("config")
	viperConfig.AddConfigPath(path)
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
	err := viperConfig.BindPFlag(key, flag)
	if err != nil {
		panic(err)
	}
}

func GetFlag(key string) any {
	return viperConfig.Get(key)
}
