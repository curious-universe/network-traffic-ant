package collect

import (
	"fmt"
	"github.com/curious-universe/network-traffic-ant/elasticsearch"
	"github.com/google/gopacket"
	"strings"
)

func SavePacketInfo(packet gopacket.Packet) {
	applicationLayer := packet.ApplicationLayer()
	if applicationLayer != nil {
		fmt.Println("Application layer/Payload found.")
		payload := string(applicationLayer.Payload())
		fmt.Printf("%q\n", payload)
		// Search for a string inside the payload
		//if strings.Contains(payload, "HTTP") {
		//	fmt.Println("HTTP found!")
		//}
		var host = ""
		var requestMethod = "Get"
		fmt.Println(strings.Split(payload, "\r\n"))
		for _, line := range strings.Split(payload, "\r\n") {
			if strings.Contains(line, "Host") {
				host = line[len("Host: "):]
				fmt.Println(line)
			}
		}
		if host != "" {
			elasticsearch.Create("application-into-packet",
				fmt.Sprintf(`{"payload":"%s","host":"%s","request_method":"%s"}`,
					"payload",
					host,
					requestMethod,
				),
			)
		}
	}
}
