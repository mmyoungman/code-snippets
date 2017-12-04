import re

file = open("02input.txt", 'r')

list = []

while True:
  line = file.readline()
  if line == '':
    break
  line = line.rstrip('\n')
  list.append(line)

file.close()

x, y = 1, 1
for digit in list:
   for i in range(0, len(digit)):
      if digit[i] == "U":
         y += 1
      elif digit[i] == "R":
         x += 1
      elif digit[i] == "D":
         y -= 1
      else:
         x -= 1

      if x < 0:
         x = 0
      if x > 2:
         x = 2
      if y < 0:
         y = 0
      if y > 2:
         y = 2

   if x == 0 and y == 2:
      print("1")
   if x == 1 and y == 2:
      print("2")
   if x == 2 and y == 2:
      print("3")
   if x == 0 and y == 1:
      print("4")
   if x == 1 and y == 1:
      print("5")
   if x == 2 and y == 1:
      print("6")
   if x == 0 and y == 0:
      print("7")
   if x == 1 and y == 0:
      print("8")
   if x == 2 and y == 0:
      print("9")
   