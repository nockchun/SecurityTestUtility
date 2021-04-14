#! /usr/bin/env python
from scapy.all import *

eth = Ether()
arp = ARP()

eth.dst = 'ff:ff:ff:ff:ff:fE'
arp.op = 1
arp.pdst = '192.168.0.20'
pkt = srp1(eth/arp)
pkt.show()
