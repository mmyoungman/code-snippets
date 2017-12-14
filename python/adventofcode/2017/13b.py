file = open("13input.txt", 'r')

firewall = {} 

while True:
  line = file.readline()
  if line == '':
    break
  line = line.rstrip('\n')
  line = line.split()
  for i in range(len(line)):
     line[i] = line[i].rstrip(':')
  firewall[int(line[0])] = int(line[1])

file.close()

delay = 0
while True:
   caught = False
   for pos in firewall:
      if (pos+delay) % (2*(firewall[pos]-1)) == 0:
         caught = True
         break
   if caught:
      delay += 1
   else:
      break

print(delay)