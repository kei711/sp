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
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use: "remove",
	Run: func(cmd *cobra.Command, args []string) {
		selectedCommand := selectCommandPrompt()
		if selectedCommand == "" {
			os.Exit(0)
		}

		commands := viper.GetStringSlice("commands")
		var after []string
		isMatched := false
		for _, v := range commands {
			if v == selectedCommand {
				isMatched = true
				continue
			}
			after = append(after, v)
		}
		viper.Set("commands", after)

		viper.WriteConfig()

		if isMatched {
			fmt.Println("command delete. " + selectedCommand)
		}
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
