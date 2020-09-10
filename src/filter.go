package main

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

func getPacketInfo(packet gopacket.Packet) {
	if ipLayer := packet.Layer(layers.LayerTypeIPv4); ipLayer != nil {
		fmt.Println("IP Layer found")
		ip, _ := ipLayer.(*layers.IPv4)

		fmt.Println("From ", ip.SrcIP, " to ", ip.DstIP)
		fmt.Println("Protocol: ", ip.Protocol)
	}

}