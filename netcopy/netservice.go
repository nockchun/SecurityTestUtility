package main

import (
	"bufio"
	b64 "encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"regexp"
	"strconv"
	"strings"
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
	remotefile         = make(map[int]string)
	receiveName        = ""
	snapshot_len int32 = 1024
	promiscuous  bool  = false
	err          error
	timeout      time.Duration = 5 * time.Second
	handle       *pcap.Handle
)

func rx(device string, name string) error {
	receiveName = name
	// Open device
	handle, err = pcap.OpenLive(device, snapshot_len, promiscuous, timeout)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	// Set filter
	var filter string = "icmp[icmptype] == icmp-echo"
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
		icmp, _ := icmpLayer.(*layers.ICMPv4)
		re, _ := regexp.Compile(`^NETCOPY_(\d+)\/(\d+)_(.+)$`)
		res := re.FindStringSubmatch(string(icmp.Payload))
		if len(res) == 4 {
			seq, _ := strconv.Atoi(res[1])
			remotefile[seq] = res[3]
		}
		amt, _ := strconv.Atoi(res[2])
		fmt.Printf("\rReceived %d/%d", len(remotefile), amt)
		if amt == len(remotefile) {
			builder := strings.Builder{}
			for i := 1; i <= amt; i++ {
				builder.WriteString(remotefile[i])
			}
			buff, _ := b64.StdEncoding.DecodeString(builder.String())
			ioutil.WriteFile(receiveName, buff, 07440)
			os.Exit(0)
		}
	}
}

func tx(file string, fragments int, targetHost string) error {
	tf, _ := os.Open(file)
	reader := bufio.NewReader(tf)
	content, _ := ioutil.ReadAll(reader)
	encoded := []byte(b64.StdEncoding.EncodeToString(content))
	// fmt.Println("ENCODED: " + string(encoded))
	sendICMP(encoded, fragments, targetHost)

	return nil
}

func sendICMP(msg []byte, fragment int, targetHost string) {
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
		time.Sleep(100 * time.Microsecond)
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
		if _, err := c.WriteTo(wb, &net.IPAddr{IP: net.ParseIP(targetHost)}); err != nil {
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
