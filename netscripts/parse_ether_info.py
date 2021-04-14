from subprocess import check_output

ifinfo = check_output(['ifconfig', 'eth0'])
print(ifinfo)
print("------------------------------------")

ifinfos = ifinfo.split()
print(ifinfos)
print("------------------------------------")

iface, ipv4, mac, bcast, nmask, ipv6 = (ifinfos[i] for i in (0, 6, 4, 7, 8, 11))
print("iface : %s" % iface)
print(f"ipv4  : {ipv4[5:]}")
print("mac   : %s" % mac)
print("bcast : %s" % bcast[6:])
print("nmask : %s" % nmask[5:])
print("ipv6  : %s" % ipv6)
