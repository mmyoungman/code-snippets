file = open("12input.txt", 'r')

list = []

while True:
  line = file.readline()
  if line == '':
    break
  line = line.rstrip('\n')
  line = line.split()
  for i in range(len(line)):
     line[i] = line[i].rstrip(',')
  list.append(line)

file.close()

pipes = []
for line in list:
   pipe = []
   for i in range(len(line)):
      if line[i].isalnum():
         pipe.append(int(line[i]))
   pipes.append(pipe)

result = set([0])
oldLen, newLen = 0, 1
while newLen != oldLen:
   oldLen = len(result)
   for pipe in pipes:
      if pipe[0] in result:
         result.update(pipe[1:])
   newLen = len(result)

print(len(result))