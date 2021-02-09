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
