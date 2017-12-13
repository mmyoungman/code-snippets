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

result = [0]
oldLen, newLen = 0, 1
while newLen != oldLen:
   oldLen = len(result)
   for pipe in pipes:
      if pipe[0] in result:
         result += pipe[1:]
         newResult = [] 
         for item in result:
            if item not in newResult:
               newResult.append(item)
         result = newResult
   newLen = len(result)

print(len(result))