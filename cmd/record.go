package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

type recordCmdArgs struct {
	Interface string
	BPF       string
}

var RecordCmdArgs recordCmdArgs

func init() {
	recordCmd.Flags().StringVarP(&RecordCmdArgs.Interface, "interface", "i", "s", "Name of network card interface")
	recordCmd.Flags().StringVarP(&RecordCmdArgs.BPF, "bpf", "b", "", "BPF filter")
	rootCmd.AddCommand(recordCmd)
}

var recordCmd = &cobra.Command{
	Use:   "record",
	Short: getRecordCmdShort(),
	Long:  getRecordCmdLong(),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%+v\n", RecordCmdArgs)
	},
}

func getRecordCmdShort() string {
	return fmt.Sprintf("This cmd is record")
}

func getRecordCmdLong() string {
	return fmt.Sprintf("This cmd is record long")
}
