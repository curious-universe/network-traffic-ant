package main

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"log"
	"time"
)

func findDevice() (devices []pcap.Interface) {
	devices, err := pcap.FindAllDevs()
	if err != nil {
		panic(err)
	}
	return
}

func main() {
	//devices := findDevice()
	devices := []pcap.Interface{{Name: "lo0"}}
	log.Printf("%+v\n", devices)
	var snapshot_len int32 = 1024
	promiscuous := false
	timeout := 100 * time.Millisecond
	handle, err := pcap.OpenLive(devices[0].Name, snapshot_len, promiscuous, timeout)
	if err != nil {
		log.Fatal(err)
	}

	handle.SetBPFFilter("port 3001")
	log.Printf("%s:%+v, %+v", devices[0].Name, handle, handle.LinkType())

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	ch := packetSource.Packets()
	for {
		select {
		case packet := <-ch:
			al := packet.ApplicationLayer()
			if al != nil {
				log.Println(packet.Metadata().CaptureInfo)
				log.Println(packet.NetworkLayer().NetworkFlow())
				log.Printf("application:\n%s\n", string(al.Payload()))
				nl := packet.NetworkLayer()
				if nl != nil {
					// IPv4 or IPv6
					log.Printf("network[layerType()]:\n%s\n", nl.LayerType())
					log.Printf("network[LayerPayload()]:\n%s\n", nl.LayerPayload())
					log.Printf("network[LayerContents()]:\n%s\n", nl.LayerContents())
					log.Printf("network[NetworkFlow()]:\n%s\n", nl.NetworkFlow())
				}
				tl := packet.TransportLayer()
				if tl != nil {
					// UDP or TCP
					log.Printf("Transport[layerType()]:\n%s\n", tl.LayerType())
					log.Printf("Transport[LayerPayload()]:\n%s\n", tl.LayerPayload())
					log.Printf("Transport[LayerContents()]:\n%s\n", tl.LayerContents())
					log.Printf("Transport[TransportFlow()]:\n%s\n", tl.TransportFlow())
				}
				ly := packet.LinkLayer()
				if ly != nil {
					log.Printf("Link[layerType()]:\n%s\n", ly.LayerType())
					log.Printf("Link[LayerPayload()]:\n%s\n", ly.LayerPayload())
					log.Printf("Link[LayerContents()]:\n%s\n", ly.LayerContents())
				}
				el := packet.ErrorLayer()
				if el != nil {
					log.Printf("Error[layerType()]:\n%s\n", el.LayerType())
					log.Printf("Error[LayerPayload()]:\n%s\n", el.LayerPayload())
					log.Printf("Error[LayerContents()]:\n%s\n", el.LayerContents())
				}
			}
		}
	}
}
