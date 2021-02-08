package cmd

import (
	"fmt"
	"github.com/whiteCcinn/network-traffic-ant/constant"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(recordCmd)
}

var recordCmd = &cobra.Command{
	Use:   "record",
	Short: getVersionCmdShort(),
	Long:  getVersionCmdLong(),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Version is v%s\n", constant.VERSION)
		fmt.Printf("BuildTS is %s\n", constant.BuildTS)
		fmt.Printf("GitHash is %s\n", constant.GitHash)
		fmt.Printf("GitBranch is %s\n", constant.GitBranch)
	},
}

func getRecordCmdShort() string {
	return fmt.Sprintf("This cmd is record")
}

func getRecordCmdLong() string {
	return fmt.Sprintf("This cmd is record")
}
