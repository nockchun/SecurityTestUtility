#! /usr/bin/env python

import xml.etree.ElementTree as etree
import csv

tree = etree.parse("result.xml")
root = tree.getroot()
hosts = root.findall("host")

csv_title = ["state", "addr"]
csv_rows = []
for host in hosts:
    row = []
    row.append(host.findall("status")[0].attrib["state"])
    row.append(host.findall("address")[0].attrib["addr"])
    csv_rows.append(row)

csv_file = open("nmap.csv", "w")
csv_writer = csv.writer(csv_file)
csv_writer.writerow(csv_title)
csv_writer.writerows(csv_rows)
csv_file.close()