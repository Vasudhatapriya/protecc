package filter

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	// "github.com/SaurusXI/filter/ipfilter"
)

func getPacketInfo(packet gopacket.Packet) {
	if ipLayer := packet.Layer(layers.LayerTypeIPv4); ipLayer != nil {
		fmt.Println("IP Layer found")
		ip, _ := ipLayer.(*layers.IPv4)

		fmt.Println("From ", ip.SrcIP, " to ", ip.DstIP)
		fmt.Println("Protocol: ", ip.Protocol)
	}

	if tcpLayer := packet.Layer(layers.LayerTypeTCP); tcpLayer != nil {
		fmt.Println("TCP Layer found")
		tcp, _ := tcpLayer.(*layers.TCP)

		fmt.Println("From port ", tcp.SrcPort, " to ", tcp.DstPort)
		fmt.Println("Sequence number: ", tcp.Seq)
	}

	if err := packet.ErrorLayer(); err != nil {
		fmt.Println("Error decoding packet:", err)
	}
}