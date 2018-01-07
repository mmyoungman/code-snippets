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

def maxstrength(bridge, port, length):
   matches = []
   # loop backwards so indexes aren't screwed by "del newbridge[blah]"
   for i in range(len(bridge)-1, -1, -1): 
      if bridge[i][0] == port:
         matches.append(bridge[i] + [0, i])
      elif bridge[i][1] == port:
         matches.append(bridge[i] + [1, i])
   
   if len(matches) < 1:
      return 0, length
   else:
      max = 0 
      maxlen = 0
      for match in matches:
         newport = 0
         if match[2] == 0:
            newport = 1
         newbridge = bridge[:]
         del newbridge[match[3]]
         newvalue, newlength = maxstrength(newbridge, match[newport], length+1)
         tempvalue = match[0] + match[1] + newvalue
         if newlength > maxlen or (newlength == maxlen and tempvalue > max):
            max = tempvalue
            maxlen = newlength
      return max, maxlen

print(maxstrength(pipes, 0, 0)[0])