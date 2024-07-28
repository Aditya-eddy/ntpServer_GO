package main

import (
	"fmt"
	"net"

	"github.com/Aditya-eddy/ntp/packet"
)

func main() {
	packetData := packet.NewNTPPacket()
	data, err := packetData.MarshalBinary()

	if err != nil {
		fmt.Println("Error in marshalling the packet: ", err)
		return
	}

	conn, err := net.Dial("udp", "129.6.15.28:123")

	if err != nil {
		fmt.Println("Error in connecting to the NTP Server: ", err)
		return
	}

	defer conn.Close()

	_, err = conn.Write(data)

	if err != nil {
		fmt.Println("Error in sending the NTP Packet: ", err)
		return
	}

	buffer := make([]byte, 48)

	_, err = conn.Read(buffer)

	if err != nil {
		fmt.Println("Error in receiving the NTP Packet: ", err)
		return
	}

	ntpResponse := &packet.NTPPACKET{}

	err = ntpResponse.UnMarshalBinary(buffer)

	if err != nil {
		fmt.Println("Error in unmarshalling the buffer received ", err)
	}

	fmt.Printf("Received NTP Response: %x\n", ntpResponse)

	packet.DecodeNTPPacket(ntpResponse)

}
