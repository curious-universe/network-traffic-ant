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
	ps "github.com/curious-universe/go-ps"
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
		procs, _ := ps.Processes()

		for _, p := range procs {
			zaplog.S().Infof("%#v", p)
			zaplog.S().Infof("%#v", p.Executable())
		}
		zaplog.S().Infof("%#v", procs)
	},
}

func getRecordCmdShort() string {
	return fmt.Sprintf("This cmd is begining record")
}

func getRecordCmdLong() string {
	return fmt.Sprintf("This cmd is begining record, this will push the network flow to the target")
}
