import sys

file = open("20input.txt", 'r')

list = []

while True:
  line = file.readline()
  if line == '':
    break
  line = line.rstrip('\n')
  list.append(line.split(', '))

file.close()

for i in range(len(list)):
   for j in range(len(list[i])):
      list[i][j] = list[i][j][3:].rstrip('>').split(',')
   
for i in range(len(list)):
   for j in range(len(list[i])):
      for k in range(len(list[i][j])):
         list[i][j][k] = int(list[i][j][k])

def updateV(p):
   result = [0,0,0]
   for i in range(3):
      result[i] = p[1][i] + p[2][i]
   return result
def updateP(p):
   result = [0,0,0]
   for i in range(3):
      result[i] = p[0][i] + p[1][i]
   return result
def getDist(p):
   return abs(p[0][0]) + abs(p[0][1]) + abs(p[0][2])

for i in range(1000):
   for j in range(len(list)):
      list[j][1] = updateV(list[j])
      list[j][0] = updateP(list[j])

minIndex = -1
minDist = 100000000000000
for i in range(len(list)):
   if getDist(list[i]) < minDist:
      minDist = getDist(list[i])
      minIndex = i

print(minIndex)
