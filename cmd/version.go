/*
Copyright © 2021 curious-universe

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmd

import (
	"fmt"
	"github.com/curious-universe/network-traffic-ant/constant"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: getVersionCmdShort(),
	Long:  getVersionCmdLong(),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Version is v%s\n", constant.VERSION)
		fmt.Printf("BuildTS is %s\n", constant.BuildTS)
		fmt.Printf("GitHash is %s\n", constant.GitHash)
		fmt.Printf("GitBranch is %s\n", constant.GitBranch)
	},
}

func getVersionCmdShort() string {
	return fmt.Sprintf("Print the version number of %s", constant.AppName)
}

func getVersionCmdLong() string {
	return fmt.Sprintf("All software has versions. This is %s's", constant.AppName)
}
