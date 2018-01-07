import itertools
file = open("24input.txt", 'r')

pipes = []

while True:
   line = file.readline()
   if line == '':
      break
   pipes.append(line.rstrip('\n').split('/'))

file.close()

for i in range(len(pipes)):
   pipes[i][0] = int(pipes[i][0])
   pipes[i][1] = int(pipes[i][1])

def maxstrength(bridge, port):
   matches = []
   for i in range(len(bridge)-1, -1, -1):
      if bridge[i][0] == port:
         matches.append(bridge[i] + [0, i])
      elif bridge[i][1] == port:
         matches.append(bridge[i] + [1, i])
   
   if len(matches) < 1:
      return 0
   else:
      max = 0 
      for match in matches:
         newport = 0
         if match[2] == 0:
            newport = 1
         newbridge = bridge[:]
         del newbridge[match[3]]
         temp = match[0] + match[1] + maxstrength(newbridge, match[newport])
         if temp > max:
            max = temp
      return max

print(maxstrength(pipes, 0))