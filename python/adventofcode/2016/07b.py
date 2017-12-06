import re

file = open("07input.txt", 'r')

list = []

while True:
  line = file.readline()
  if line == '':
    break
  line = line.rstrip('\n')
  list.append(line)

file.close()

numssl = 0

for ipaddr in list:
   insidebrac = False
   abas = []
   babs = []
   for i in range(0, len(ipaddr)-2):
      if "[" in ipaddr[i:i+3]:
         insidebrac = True
         continue
      if "]" in ipaddr[i:i+3]:
         insidebrac = False
         continue
      if not insidebrac and ipaddr[i] != ipaddr[i+1] and ipaddr[i] == ipaddr[i+2]:
         abas.append(ipaddr[i:i+3])
      if insidebrac and ipaddr[i] != ipaddr[i+1] and ipaddr[i] == ipaddr[i+2]:
         babs.append(ipaddr[i:i+3])

   abababmatch = False
   for aba in abas:
      for bab in babs:
         if not abababmatch and aba[0] == bab[1] and aba[1] == bab[0]:
            numssl += 1
            abababmatch = True

print(numssl)