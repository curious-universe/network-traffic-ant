package main

import "github.com/google/gopacket/pcap"

func findDevice() (devices []pcap.Interface) {
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
