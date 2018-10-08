// Copyright © 2018 kei711 <kei.yam.711@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"github.com/kei711/sp/cmd/runner"
	"github.com/kei711/sp/config"
	"github.com/kei711/sp/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var cfgFile string
var verbose bool

var rootCmd = &cobra.Command{
	Use: "sp",
	Run: func(cmd *cobra.Command, args []string) {
		selectedCommand := selectCommandPrompt()
		if selectedCommand == "" {
			os.Exit(0)
		}
		runner.Run(selectedCommand, verbose)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.sp.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
}

func initConfig() {
	config.Init()

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home := util.GetHomeDir()

		viper.AddConfigPath(home)
		viper.SetConfigName(".sp")

		viper.SetConfigFile(home + "/" + ".sp.yaml")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		viper.WriteConfig()
	} else if verbose {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func selectCommandPrompt() string {
	selectedCommand := prompt.Input(">>> ", commandsCompleter,
		prompt.OptionTitle("choose command"),
		prompt.OptionPrefixTextColor(prompt.Blue),
		prompt.OptionSelectedDescriptionBGColor(prompt.LightGray),
		prompt.OptionSelectedSuggestionBGColor(prompt.LightGray),
		prompt.OptionSelectedDescriptionTextColor(prompt.DarkGray),
		prompt.OptionSelectedSuggestionTextColor(prompt.DarkGray),
	)

	return selectedCommand
}

func commandsCompleter(d prompt.Document) []prompt.Suggest {
	var commandSuggest []prompt.Suggest
	for _, command := range config.GetCommands() {
		commandSuggest = append(commandSuggest, prompt.Suggest{Text: command})
	}

	return util.FilterFuzzy(commandSuggest, d.GetWordBeforeCursor())
}
