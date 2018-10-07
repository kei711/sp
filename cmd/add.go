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
	"github.com/kei711/symfony-console-commands-prompt/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"path/filepath"
	"strings"
)

var commandPath string

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use: "add",
	Run: func(cmd *cobra.Command, args []string) {
		commandPath = strings.Replace(commandPath, "~", util.GetHomeDir(), 1)
		commandPath, _ := filepath.Abs(commandPath)

		if !util.FileExists(commandPath) {
			fmt.Println("command not found.")
			return
		}

		commands := viper.GetStringSlice("commands")
		commands = append(commands, commandPath)
		viper.Set("commands", commands)

		viper.WriteConfig()

		fmt.Println("command added. " + commandPath)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringVarP(&commandPath, "command", "c", "", "command path")
}
