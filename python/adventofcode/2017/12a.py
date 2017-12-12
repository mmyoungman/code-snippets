file = open("12input.txt", 'r')

pipes = []
pipeDict = {}

while True:
   line = file.readline()
   if line == '':
      break
   line = line.rstrip('\n').split()
   for i in range(len(line)):
      line[i] = line[i].rstrip(',')
   pipeDict[line[0]] = line[2:]
   pipes.append(line[0])

file.close()

# print(pipeDict)

group = ['0']

for i in range(30):
   for pipe in pipes:
      if pipe in group:
         group += pipeDict[pipe]
      newgroup = []
      for j in group:
         if j not in newgroup:
            newgroup.append(j)
      group = newgroup

print(group)
print(len(group))