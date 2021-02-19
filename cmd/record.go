package cmd

import (
	"fmt"
	"github.com/curious-universe/network-traffic-ant/config"
	"github.com/curious-universe/network-traffic-ant/zaplog"
	"github.com/spf13/cobra"
)

type recordCmdArgs struct {
	Interface   string
	BPF         string
	ProcessName string
}

var RecordCmdArgs recordCmdArgs

func init() {
	recordCmd.Flags().StringVarP(&RecordCmdArgs.Interface, "interface", "i", "lo", "Name of network card interface")
	recordCmd.Flags().StringVarP(&RecordCmdArgs.BPF, "bpf", "b", "", "BPF filter")
	recordCmd.Flags().StringVarP(&RecordCmdArgs.ProcessName, "process", "p", "", "The Process Name")

	if err := recordCmd.MarkFlagRequired("process"); err != nil {

	}

	rootCmd.AddCommand(recordCmd)
}

var recordCmd = &cobra.Command{
	Use:   "record",
	Short: getRecordCmdShort(),
	Long:  getRecordCmdLong(),
	Run: func(cmd *cobra.Command, args []string) {
		zaplog.S().Infof("%#v", RecordCmdArgs)
		zaplog.S().Infof("%#v", config.GetGlobalConfig())
	},
}

func getRecordCmdShort() string {
	return fmt.Sprintf("This cmd is begining record")
}

func getRecordCmdLong() string {
	return fmt.Sprintf("This cmd is begining record, this will push the network flow to the target")
}
