import re

file = open("05input.txt", 'r')

list = []

while True:
  line = file.readline()
  if line == '':
    break
  line = line.rstrip('\n')
  list.append(int(line))

file.close()

position = 0
step = 0

while position < len(list):
   jump = list[position]
   if list[position] >= 3:
      list[position] -= 1
   else:
      list[position] += 1
   position += jump
   step += 1

print(step)