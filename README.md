# NTP Client in Go

This project implements a simple NTP client in Go. The client sends a request to an NTP server, receives the response, and decodes the NTP packet to display human-readable information.

## Files

- `main.go`: The main entry point of the application.
- `header.go`: Contains the `NTPPACKET` struct and methods for marshaling and unmarshaling NTP packets.
- `decode.go`: Contains functions to decode the NTP packet into human-readable format.

## Structs and Functions

### NTPPACKET Struct

The `NTPPACKET` struct represents an NTP packet. It includes the following fields:

- `LI_VN_MODE`: Leap Indicator, Version Number, and Mode combined into one byte.
- `Stratum`: The stratum level of the clock.
- `Poll`: The maximum interval between successive messages.
- `Precision`: The precision of the system clock.
- `RootDelay`: The total round trip delay to the primary reference source.
- `RootDispersion`: The maximum error relative to the primary reference source.
- `ReferenceID`: Identifier for the particular server or reference clock.
- `ReferenceTS`: The time at which the system clock was last set or corrected.
- `OriginTS`: The time at which the request departed the client for the server.
- `ReceiveTS`: The time at which the request arrived at the server.
- `TransmitTS`: The time at which the reply departed the server for the client.

### Functions

- `NewNTPPacket() *NTPPACKET`: Creates a new NTP packet with default values.
- `(*NTPPACKET) MarshalBinary() ([]byte, error)`: Marshals the NTP packet to binary format.
- `(*NTPPACKET) UnmarshalBinary(data []byte) error`: Unmarshals binary data into an NTP packet.
- `DecodeNTPPacket(p *NTPPACKET)`: Decodes and prints the fields of an NTP packet in human-readable format.
- `NTPTime(ntpTimestamp uint64) time.Time`: Converts an NTP timestamp to a human-readable time.

## Usage

To run the NTP client:

1. Clone the repository.
2. Ensure you have Go installed.
3. Run `go run main.go`.

The client will send a request to the NTP server, receive the response, and print the decoded NTP packet fields.

## Example Output

```plaintext
Received NTP Response: &{LI_VN_MODE:0x1c Stratum:1 Poll:10 Precision:-29 RootDelay:0x20 RootDispersion:0x4e495354 ReferenceID:0xea50f100 ReferenceTS:0xea50f15000000000 OriginTS:0xea50f15000000000 ReceiveTS:0xea50f150ee228d10 TransmitTS:0xea50f150ee22a4b0}
Leap Indicator: 0
Version Number: 3
Mode: 4
Stratum: 1
Poll: 10
Precision: -29
Root Delay: 0
Root Dispersion: 0
Reference ID: ea50f100
Reference Timestamp: 2024-07-28 10:24:00 +0000 UTC
Originate Timestamp: 2024-07-28 10:24:00 +0000 UTC
Receive Timestamp: 2024-07-28 10:24:01.92959356 +0000 UTC
Transmit Timestamp: 2024-07-28 10:24:01.92961584 +0000 UTC
```
