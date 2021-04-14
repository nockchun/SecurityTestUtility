#! /usr/bin/env python
from scapy.all import *
from scapy.layers.inet import *

src="10.0.24.7"
pcap_file = "/home/pent/dump"
pcap = rdpcap(pcap_file)

data = ""
myData = ""
for packet in pcap:
    if IP in packet and TCP in packet:
        data += packet[IP].summary() + "\n"
        if packet[IP].src == src:
            myData += packet[IP].summary() + "\n"


fData = open("/home/pent/result.dat", 'w')
fData.write(data)
fData.close()

fMyData = open("/home/pent/result_my.dat", 'w')
fMyData.write(myData)
fMyData.close()
