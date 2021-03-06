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
	"github.com/curious-universe/network-traffic-ant/collect"
	"github.com/curious-universe/network-traffic-ant/config"
	"github.com/curious-universe/network-traffic-ant/elasticsearch"
	"github.com/curious-universe/network-traffic-ant/nerror"
	"github.com/curious-universe/network-traffic-ant/process"
	"github.com/curious-universe/network-traffic-ant/zaplog"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/spf13/cobra"
	"strconv"
	"strings"
	"time"
)

type collectCmdArgs struct {
	Interface     string
	BPF           string
	ProcessBinary string
	ProcessPid    int
}

var CollectCmdArgs collectCmdArgs

func init() {
	recordCmd.Flags().StringVarP(&CollectCmdArgs.Interface, "interface", "i", findFirstDevice().Name, "Name of network card interface")
	recordCmd.Flags().StringVarP(&CollectCmdArgs.BPF, "bpf", "b", "", "BPF filter")
	recordCmd.Flags().StringVarP(&CollectCmdArgs.ProcessBinary, "process_binary", "p", "", "The Process Name")
	recordCmd.Flags().IntVarP(&CollectCmdArgs.ProcessPid, "pid", "", 0, "The Process pid")

	if err := recordCmd.MarkFlagRequired("process_binary"); err != nil {
		zaplog.S().Fatal(err)
	}
	rootCmd.AddCommand(recordCmd)
}

var recordCmd = &cobra.Command{
	Use:   "collect",
	Short: getRecordCmdShort(),
	Long:  getRecordCmdLong(),
	Run: func(cmd *cobra.Command, args []string) {
		zaplog.S().Infof("%#v", config.GetGlobalConfig())
		// Find Process
		_, err := process.FindProcessByName(CollectCmdArgs.ProcessBinary)
		if err == nerror.ErrTooManySameNameProcess {
			if CollectCmdArgs.ProcessPid == 0 {
				zaplog.S().Fatal(nerror.ErrProcessPidMustNotNil.Error())
			}
			_, err = process.FindProcessByNameAndPid(CollectCmdArgs.ProcessBinary, CollectCmdArgs.ProcessPid)
			if err == nerror.ErrNotFoundProcess {
				zaplog.S().Fatal(nerror.ErrNotFoundProcess.Error())
			}
		}
		nerror.MustNil(err)

		// Open device
		handle, err := pcap.OpenLive(CollectCmdArgs.Interface, 1024, false, 30*time.Second)
		nerror.MustNil(err)
		defer handle.Close()
		packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
		for packet := range packetSource.Packets() {
			//printPacketInfo(packet)
			collect.SavePacketInfo(packet)
		}
	},
}

func getRecordCmdShort() string {
	return fmt.Sprintf("This cmd is begining record")
}

func getRecordCmdLong() string {
	return fmt.Sprintf("This cmd is begining record, this will push the network flow to the target")
}

var printCnt = 1

func printPacketInfo(packet gopacket.Packet) {
	fmt.Println("=============printPacketInfo " + strconv.Itoa(printCnt) + "=============")
	// Let's see if the packet is an ethernet packet
	ethernetLayer := packet.Layer(layers.LayerTypeEthernet)
	if ethernetLayer != nil {
		fmt.Println("Ethernet layer detected.")
		ethernetPacket, _ := ethernetLayer.(*layers.Ethernet)
		fmt.Println("Source MAC: ", ethernetPacket.SrcMAC)
		fmt.Println("Destination MAC: ", ethernetPacket.DstMAC)
		// Ethernet type is typically IPv4 but could be ARP or other
		fmt.Println("Ethernet type: ", ethernetPacket.EthernetType)
		fmt.Println()
		elasticsearch.Create("ethernet-packet",
			fmt.Sprintf(`{"source_mac":"%s","dst_mac":"%s","ethernet_type":"%s"}`,
				ethernetPacket.SrcMAC.String(),
				ethernetPacket.DstMAC.String(),
				ethernetPacket.EthernetType.String()))
	}
	// Let's see if the packet is IP (even though the ether type told us)
	ipLayer := packet.Layer(layers.LayerTypeIPv4)
	if ipLayer != nil {
		fmt.Println("IPv4 layer detected.")
		ip, _ := ipLayer.(*layers.IPv4)
		// IP layer variables:
		// Version (Either 4 or 6)
		// IHL (IP Header Length in 32-bit words)
		// TOS, Length, Id, Flags, FragOffset, TTL, Protocol (TCP?),
		// Checksum, SrcIP, DstIP
		fmt.Printf("From %s to %s\n", ip.SrcIP, ip.DstIP)
		fmt.Println("Protocol: ", ip.Protocol)
		fmt.Println()
		elasticsearch.Create("ip-packet", fmt.Sprintf(`{"source_ip":"%s","dst_ip":"%s","protocol":"%s"}`,
			ip.SrcIP.String(),
			ip.DstIP.String(),
			ip.Protocol.String()))
	}
	// Let's see if the packet is TCP
	tcpLayer := packet.Layer(layers.LayerTypeTCP)
	if tcpLayer != nil {
		fmt.Println("TCP layer detected.")
		tcp, _ := tcpLayer.(*layers.TCP)
		// TCP layer variables:
		// SrcPort, DstPort, Seq, Ack, DataOffset, Window, Checksum, Urgent
		// Bool flags: FIN, SYN, RST, PSH, ACK, URG, ECE, CWR, NS
		fmt.Printf("From port %d to %d\n", tcp.SrcPort, tcp.DstPort)
		fmt.Println("Sequence number: ", tcp.Seq)
		fmt.Println("Ack number: ", tcp.Ack)
		fmt.Println("SYN : ", tcp.SYN)
		fmt.Println("ACK : ", tcp.ACK)
		fmt.Println("PSH : ", tcp.PSH)
		fmt.Println("FIN : ", tcp.FIN)
		fmt.Println("RST : ", tcp.RST)
		fmt.Println("URG : ", tcp.URG)
		fmt.Println()
	}
	// Iterate over all layers, printing out each layer type
	fmt.Println("All packet layers:")
	for _, layer := range packet.Layers() {
		fmt.Println("- ", layer.LayerType())
	}
	// When iterating through packet.Layers() above,
	// if it lists Payload layer then that is the same as
	// this applicationLayer. applicationLayer contains the payload
	applicationLayer := packet.ApplicationLayer()
	if applicationLayer != nil {
		fmt.Println("Application layer/Payload found.")
		fmt.Printf("%s\n", applicationLayer.Payload())
		// Search for a string inside the payload
		if strings.Contains(string(applicationLayer.Payload()), "HTTP") {
			fmt.Println("HTTP found!")
		}
	}
	// Check for errors
	if err := packet.ErrorLayer(); err != nil {
		fmt.Println("Error decoding some part of the packet:", err)
	}
	printCnt++
}

func findAllDevice() (devices []pcap.Interface) {
	devices, err := pcap.FindAllDevs()
	if err != nil {
		panic(err)
	}
	return
}

func findFirstDevice() pcap.Interface {
	devices, err := pcap.FindAllDevs()
	if err != nil {
		panic(err)
	}
	return devices[0]
}
