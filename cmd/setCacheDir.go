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
	"github.com/kei711/sp/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"strings"
)

// setCacheDirCmd represents the setCacheDir command
var setCacheDirCmd = &cobra.Command{
	Use: "setCacheDir",
	Run: func(cmd *cobra.Command, args []string) {
		cacheDir = strings.Replace(cacheDir, "~", util.GetHomeDir(), 1)
		dir, _ := filepath.Abs(cacheDir)

		viper.Set("CacheDir", dir)
		if err := viper.WriteConfig(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

var cacheDir string

func init() {
	rootCmd.AddCommand(setCacheDirCmd)

	setCacheDirCmd.Flags().StringVarP(&cacheDir, "cache-dir", "c", "", "cache directory")
}
