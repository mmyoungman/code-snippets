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

numips = 0

for ipaddr in list:
   insidebrac = False
   abbaoutside = False
   abbainside = False
   for i in range(0, len(ipaddr)-3):
      if "[" in ipaddr[i:i+4]:
         insidebrac = True
         continue
      if "]" in ipaddr[i:i+4]:
         insidebrac = False
         continue
      if not insidebrac and ipaddr[i] != ipaddr[i+1] and ipaddr[i] == ipaddr[i+3] and ipaddr[i+1] == ipaddr[i+2]:
         abbaoutside = True
      if insidebrac and ipaddr[i] != ipaddr[i+1] and ipaddr[i] == ipaddr[i+3] and ipaddr[i+1] == ipaddr[i+2]:
         abbainside = True
   if abbainside == False and abbaoutside == True:
      numips += 1

print(numips)