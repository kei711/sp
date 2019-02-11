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
	"bufio"
	"fmt"
	"github.com/c-bata/go-prompt"
	"github.com/kei711/sp/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use: "add",
	Run: func(cmd *cobra.Command, args []string) {
		selectedCommand := prompt.Input(">>> ",
			func(d prompt.Document) []prompt.Suggest {
				var commandSuggest []prompt.Suggest
				files, err := ioutil.ReadDir("./")
				if err != nil {
					panic(err)
				}

				for _, file := range files {
					commandSuggest = append(commandSuggest, prompt.Suggest{Text: file.Name()})
				}
				return util.FilterFuzzy(commandSuggest, d.GetWordBeforeCursor())
			},
			prompt.OptionTitle("choose command"),
			prompt.OptionPrefixTextColor(prompt.Blue),
			prompt.OptionSelectedDescriptionBGColor(prompt.LightGray),
			prompt.OptionSelectedSuggestionBGColor(prompt.LightGray),
			prompt.OptionSelectedDescriptionTextColor(prompt.DarkGray),
			prompt.OptionSelectedSuggestionTextColor(prompt.DarkGray),
		)

		if selectedCommand == "" {
			os.Exit(0)
		}

		path := strings.Replace(selectedCommand, "~", util.GetHomeDir(), 1)
		commandPath, _ := filepath.Abs(path)
		if !validateCommand(commandPath) {
			os.Exit(1)
		}

		commands := viper.GetStringSlice("commands")
		commands = append(commands, commandPath)
		viper.Set("commands", commands)

		if err := viper.WriteConfig(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("command added. " + commandPath)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func validateCommand(commandPath string) bool {
	if !util.FileExists(commandPath) {
		fmt.Println("command not found.")
		return false
	}

	fp, err := os.Open(commandPath)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer fp.Close()

	// validating shebang
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		firstLine := scanner.Text()
		r := regexp.MustCompile(`^#!.*\s+php`)
		if !r.MatchString(firstLine) {
			fmt.Println("command is not PHP file.")
			return false
		}
		break
	}

	if err = scanner.Err(); err != nil {
		fmt.Println(err)
		return false
	}

	return true
}
