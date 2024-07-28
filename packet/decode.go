package packet

import (
	"fmt"
	"time"
)

func DecodeNTPPacket(p *NTPPACKET) {

	fmt.Printf("Reference Timestamp: %s\n", ntpTime(p.ReferenceTS))
	fmt.Printf("Originate Timestamp: %s\n", ntpTime(p.OriginTS))
	fmt.Printf("Receive Timestamp: %s\n", ntpTime(p.ReceiveTS))
	fmt.Printf("Transmit Timestamp: %s\n", ntpTime(p.TransmitTS))

}

func ntpTime(ntpTimeStamp uint64) time.Time {
	seconds := ntpTimeStamp >> 32

	faction := ntpTimeStamp & 0xFFFFFFFF

	unixTime := int64(seconds) - 2208988800

	nanoSeconds := (int64(faction) * 1e9) >> 32

	return time.Unix(unixTime, nanoSeconds)
}
