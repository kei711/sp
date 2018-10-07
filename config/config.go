package config

import (
	"github.com/kei711/symfony-console-commands-prompt/util"
	"github.com/spf13/viper"
)

type Config struct {
	CacheDir string
	Commands []string
}

const (
	keyCommands = "commands"
	keyCacheDir = "cacheDir"

	defaultCacheDir = ".sp"
)

func Init() {
	viper.SetConfigType("yaml")
	viper.SetDefault(keyCacheDir, util.GetHomeDir()+"/"+defaultCacheDir)
	viper.SetDefault(keyCommands, make([]string, 0, 1))
viper.ReadInConfig()
	viper.WriteConfig()
}

func GetCommands() []string {
	return viper.GetStringSlice(keyCommands)
}

func GetCacheDir() string {
	return viper.GetString(keyCacheDir)
}
