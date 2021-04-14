#!/usr/bin/env python
from scapy.all import *
from scapy.layers.inet import *

eth = Ether(dst="ff:ff:ff:ff:ff:ff")
arp = ARP(op=1, pdst="10.0.24.9")

pkt = srp1(eth/arp)
pkt.show()
