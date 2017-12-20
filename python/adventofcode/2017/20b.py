import sys

file = open("20input.txt", 'r')

prts = []

while True:
  line = file.readline()
  if line == '':
    break
  line = line.rstrip('\n')
  prts.append(line.split(', '))

file.close()

for i in range(len(prts)):
   for j in range(len(prts[i])):
      prts[i][j] = prts[i][j][3:].rstrip('>').split(',')
   
for i in range(len(prts)):
   for j in range(len(prts[i])):
      for k in range(len(prts[i][j])):
         prts[i][j][k] = int(prts[i][j][k])

for i in range(1000):
   particles = {}
   toDel = []
   for j in range(len(prts)):
      if (prts[j][0][0], prts[j][0][1], prts[j][0][2]) in particles.values():
         toDel.append(j)
         for key, value in particles.items():
            if value == (prts[j][0][0], prts[j][0][1], prts[j][0][2]):
               toDel.append(key)
      else:
         particles[j] = (prts[j][0][0], prts[j][0][1], prts[j][0][2])
   
   toDel = set(toDel)
   toDel = sorted(list(toDel), reverse=True)

   for index in toDel:
      del prts[index]

   for j in range(len(prts)):
      result = [0,0,0]
      for i in range(3):
         result[i] = prts[j][1][i] + prts[j][2][i]
      prts[j][1] = result
      result = [0,0,0]
      for i in range(3):
         result[i] = prts[j][0][i] + prts[j][1][i]
      prts[j][0] = result

print(len(prts))
