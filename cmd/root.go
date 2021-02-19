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
	"github.com/curious-universe/network-traffic-ant/config"
	"github.com/spf13/cobra"
	"os"
)

var cfgFile string
var cfgCheck bool

var rootCmd = &cobra.Command{
	Use:   "network_traffic_ant",
	Short: "A network traffic movers",
	Long:  `This tool enables process-based traffic recording and playback`,
}

var checkConfigCmd = &cobra.Command{
	Use:   "check_config",
	Short: "Config checker",
	Long:  "Check the config and path is valid",
	Run: func(cmd *cobra.Command, args []string) {
		cfgCheck = true
		initialize()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initialize)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "config.yml", "config file")

	rootCmd.AddCommand(checkConfigCmd)
}

func initialize() {
	config.InitializeConfig(cfgFile, cfgCheck)
}
