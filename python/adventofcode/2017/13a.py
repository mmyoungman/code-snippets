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

severity = 0
for pos in firewall:
   if pos % (2*(firewall[pos]-1)) == 0:
      severity += pos * firewall[pos]

print(severity)