import re

file = open("03input.txt", 'r')

list = []

while True:
  line = file.readline()
  if line == '':
    break
  list.append(line.split())

file.close()

num_triangles = 0

for tri in list:
   for i in range(len(tri)):
      tri[i] = int(tri[i])
   if tri[1] + tri[2] <= tri[0]:
      continue
   if tri[0] + tri[2] <= tri[1]:
      continue
   if tri[0] + tri[1] <= tri[2]:
      continue
   num_triangles += 1

print(num_triangles)