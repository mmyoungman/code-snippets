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

for i in range(0, len(list), 3):
   for j in range(3):
      if list[i][j] + list[i+1][j] <= list[i+2][j]:
         continue
      if list[i][j] + list[i+2][j] <= list[i+1][j]:
         continue
      if list[i+1][j] + list[i+2][j] <= list[i][j]:
         continue
      num_triangles += 1

print(num_triangles)