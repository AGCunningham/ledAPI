package wolPower

import (
	"net"
	"bytes"
	"encoding/binary"
	"fmt"
)

// This function accepts a MAC address string, and s
// Function to send a magic packet to a given mac address
func SendMagicPacket(macAddr string) error {
    magicPacket, err := NewMagicPacket(macAddr)
    if err != nil {
        return err
    }

    // Fill our byte buffer with the bytes in our MagicPacket
    var buf bytes.Buffer
    binary.Write(&buf, binary.BigEndian, magicPacket)

    // Get a UDPAddr to send the broadcast to
    udpAddr, err := net.ResolveUDPAddr("udp", "255.255.255.255:9")
    if err != nil {
        fmt.Printf("Unable to get a UDP address for %s\n", "255.255.255.255:9")
        return err
    }

    // Open a UDP connection, and defer its cleanup
    connection, err := net.DialUDP("udp", nil, udpAddr)
    if err != nil {
        fmt.Printf("Unable to dial UDP address for %s\n", "255.255.255.255:9")
        return err
    }
    defer connection.Close()

    // Write the bytes of the MagicPacket to the connection
    bytesWritten, err := connection.Write(buf.Bytes())
    if err != nil {
        fmt.Printf("Unable to write packet to connection\n")
        return err
    } else if bytesWritten != 102 {
        fmt.Printf("Warning: %d bytes written, %d expected!\n", bytesWritten, 102)
    }

    return nil
}