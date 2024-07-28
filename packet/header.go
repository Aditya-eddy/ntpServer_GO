package packet

import (
	"bytes"
	"encoding/binary"
	"time"
)

const (
	ntpEpochOffset = 2208988800
)

type NTPPACKET struct {
	LI_VN_MODE     uint8
	Stratum        uint8
	Poll           int8
	Precision      int8
	RootDelay      uint32
	RootDispersion uint32
	ReferenceID    uint32
	ReferenceTS    uint64
	OriginTS       uint64
	ReceiveTS      uint64
	TransmitTS     uint64
}

func NewNTPPacket() *NTPPACKET {
	return &NTPPACKET{
		LI_VN_MODE:     0x1b,
		Stratum:        0,
		Poll:           4,
		Precision:      -20,
		RootDelay:      0,
		RootDispersion: 0,
		ReferenceID:    0,
		ReferenceTS:    0,
		OriginTS:       0,
		ReceiveTS:      0,
		TransmitTS:     uint64(time.Now().Unix()+ntpEpochOffset) << 32,
	}
}

func (p *NTPPACKET) MarshalBinary() ([]byte, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, p)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil

}

func (p *NTPPACKET) UnMarshalBinary(data []byte) error {
	buf := bytes.NewReader(data)
	return binary.Read(buf, binary.BigEndian, p)
}
