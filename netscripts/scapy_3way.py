#!/usr/bin/python
from scapy.all import *
from scapy.layers.inet import *

ipPacket = IP(dst="www.google.com")
port = random.randrange(49152, 65535)

tcpPacket = TCP(sport = port, dport = 80, flags = "S")
syn = ipPacket/tcpPacket

synack = sr1(syn)

ack = ipPacket/TCP(sport = port, dport = 80, flags = "A")
ack[TCP].seq = synack[TCP].ack
ack[TCP].ack = synack[TCP].seq + 1
send(ack)

httpRequest = 'GET / HTTP/1.1 \r\nHost:10.2.2.1\r\n\r\n'
tcpData = TCP(sport=port, dport=80, flags="A")
tcpData.seq=synack[TCP].ack
tcpData.ack=synack[TCP].seq+1
httpPacket = ipPacket/tcpData/httpRequest
answers,unanswered = sr(httpPacket)

answers.show()
