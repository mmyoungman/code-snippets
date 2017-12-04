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

   if direction == 0:
      y += item
   elif direction == 1:
      x += item
   elif direction == 2:
      y -= item
   else:
      x -= item

print(abs(x)+abs(y))
