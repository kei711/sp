package config

import (
	"fmt"
	"github.com/kei711/sp/util"
	"github.com/spf13/viper"
	"os"
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

func Init(flags *Flags) {
	viper.AutomaticEnv()
	viper.SetDefault(keyCommands, make([]string, 0, 1))
	viper.SetDefault(keyCacheDir, util.GetHomeDir()+"/"+defaultCacheDir)

	if flags.ConfigPath != "" {
		viper.SetConfigFile(flags.ConfigPath)
	} else {
		// create .sp.yaml at home-dir
		home := util.GetHomeDir()

		viper.AddConfigPath(home)
		viper.SetConfigName(".sp")
		viper.SetConfigType("yaml")
		viper.SetConfigFile(home + "/" + ".sp.yaml")
	}

	if err := viper.ReadInConfig(); err != nil {
		if err := viper.WriteConfig(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	} else if flags.Verbose {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func HasCommands() bool {
	return viper.IsSet(keyCommands) && len(GetCommands()) > 0
}

func GetCommands() []string {
	return viper.GetStringSlice(keyCommands)
}

func GetCacheDir() string {
	return viper.GetString(keyCacheDir)
}
