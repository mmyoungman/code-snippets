import sys
import re

file = open("01input.txt", 'r')

while True:
   line = file.readline()
   if line == '':
      break
   line = line.rstrip('\n')
   list = line.split(', ')

file.close()

# 0 = N, 1 = E, 2 = S, 3 = W
direction = 0 
x, y = 0, 0
visited = []

for item in list:
   if item[0] == "R":
      direction += 1
   else:
      direction -= 1
   
   if direction == -1:
      direction = 3
   if direction == 4:
      direction = 0
   
   item = int(item[1:])

   while item > 0:
      if direction == 0:
         y += 1
      elif direction == 1:
         x += 1
      elif direction == 2:
         y -= 1
      else:
         x -= 1
      if (x, y) in visited:
         print(abs(x)+abs(y))
         sys.exit(0)
      visited.append((x,y))
      item -= 1