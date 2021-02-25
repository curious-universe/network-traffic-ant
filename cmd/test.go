/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"github.com/spf13/cobra"
	"io"
	"time"
)

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		test()
	},
}

func init() {
	rootCmd.AddCommand(testCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// testCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// testCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func test() {
	fmt.Println("test called")
	fmt.Println("findAllDevice", findAllDevice())
	fmt.Println("findFirstDevice", findFirstDevice())
	fmt.Println("createHandle", createHandle())
	packetSource, closeSource := createSource()
	defer closeSource()
	fmt.Println("createSource", packetSource)
	cnt := 0
	for {
		packet, err := packetSource.NextPacket()
		if err == io.EOF {
			return
		} else if err == nil {
			cnt++
			fmt.Println(cnt, packet)
		}
	}
}

func createSource() (*gopacket.PacketSource, func()) {
	handle := createHandle()
	return gopacket.NewPacketSource(handle, handle.LinkType()), handle.Close
}

func createHandle() *pcap.Handle {
	device := findFirstDevice()
	handle, err := pcap.OpenLive(device.Name, 1024, false, 100*time.Millisecond)
	if err != nil {
		panic(err)
	}
	return handle
}
