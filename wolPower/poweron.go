package wolPower

import (
	"errors"
	"net"
)

// This function accepts a MAC Address string, and returns a pointer to
// a MagicPacket object. A Magic Packet is a broadcast frame which
// contains 6 bytes of 0xFF followed by 16 repetitions of a given mac address.
func NewMagicPacket(mac string) (*MagicPacket, error) {
    var packet MagicPacket
    var macAddr MacAddress

    // We only support 6 byte MAC addresses
    if !re_MAC.MatchString(mac) {
        return nil, errors.New("MAC address " + mac + " is not valid.")
    }

    hwAddr, err := net.ParseMAC(mac)
    if err != nil {
        return nil, err
    }

    // Copy bytes from the returned HardwareAddr -> a fixed size MACAddress
    for idx := range macAddr {
        macAddr[idx] = hwAddr[idx]
    }

    // Setup the header which is 6 repetitions of 0xFF
    for idx := range packet.header {
        packet.header[idx] = 0xFF
    }

    // Setup the payload which is 16 repetitions of the MAC addr
    for idx := range packet.payload {
        packet.payload[idx] = macAddr
    }

    return &packet, nil
}