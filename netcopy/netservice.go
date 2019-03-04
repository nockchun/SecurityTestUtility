package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

func printDevice() error {
	// Find all devices
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatal(err)
		return err
	}

	// Print device information
	fmt.Println("Devices found:")
	for _, device := range devices {
		fmt.Println("\nName: ", device.Name)
		fmt.Println("Description: ", device.Description)
		fmt.Println("Devices addresses: ")
		for _, address := range device.Addresses {
			fmt.Println("- IP address: ", address.IP)
			fmt.Println("- Subnet mask: ", address.Netmask)
		}
	}

	return nil
}

var (
	snapshot_len int32 = 1024
	promiscuous  bool  = false
	err          error
	timeout      time.Duration = 10 * time.Second
	handle       *pcap.Handle
)

func rx(device string) error {
	// Open device
	handle, err = pcap.OpenLive(device, snapshot_len, promiscuous, timeout)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	// Set filter
	var filter string = "icmp"
	err = handle.SetBPFFilter(filter)
	if err != nil {
		log.Fatal(err)
	}

	// Use the handle as a packet source to process all packets
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		assembleICMP(packet)
	}

	return nil
}

func assembleICMP(packet gopacket.Packet) {
	icmpLayer := packet.Layer(layers.LayerTypeICMPv4)
	if icmpLayer != nil {
		// fmt.Println("ICMP layer detected.")
		icmp, _ := icmpLayer.(*layers.ICMPv4)
		if icmp.TypeCode.Code() == 8 {
			fmt.Println("payload: ", string(icmp.Payload))
		}
	}
}

func tx(file string, fragments int) error {
	tf, _ := os.Open(file)
	reader := bufio.NewReader(tf)
	content, _ := ioutil.ReadAll(reader)
	encoded := []byte(base64.StdEncoding.EncodeToString(content))
	// fmt.Println("ENCODED: " + string(encoded))
	sendICMP(encoded, fragments)

	return nil
}

func sendICMP(msg []byte, fragment int) {
	targetIP := "10.0.2.5"
	c, err := icmp.ListenPacket("ip4:icmp", "0.0.0.0")
	if err != nil {
		log.Fatalf("listen err, %s", err)
	}
	defer c.Close()

	var msgC int = 1
	var msgT int = len(msg)
	var msgS int = 0
	var msgE int = fragment
	var msgA int = (msgT / msgE) + 1
	if msgE > msgT {
		msgE = msgT
	}
	for msgE <= msgT {
		wm := icmp.Message{
			Type: ipv4.ICMPTypeEcho, Code: 0,
			Body: &icmp.Echo{
				ID: os.Getpid() & 0xffff, Seq: 1,
				Data: []byte("NETCOPY_" + strconv.Itoa(msgC) + "/" + strconv.Itoa(msgA) + "_" + string(msg[msgS:msgE])),
			},
		}
		wb, err := wm.Marshal(nil)
		if err != nil {
			log.Fatal(err)
		}
		if _, err := c.WriteTo(wb, &net.IPAddr{IP: net.ParseIP(targetIP)}); err != nil {
			log.Fatalf("WriteTo err, %s", err)
		}

		if msgE == msgT {
			break
		}
		msgC += 1
		msgS = msgE
		msgE += fragment
		if msgE > msgT {
			msgE = msgT
		}
	}
}
