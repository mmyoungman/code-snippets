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

keypad = {}
keypad[(2, 4)] = "1"
keypad[(1, 3)] = "2"
keypad[(2, 3)] = "3"
keypad[(3, 3)] = "4"
keypad[(0, 2)] = "5"
keypad[(1, 2)] = "6"
keypad[(2, 2)] = "7"
keypad[(3, 2)] = "8"
keypad[(4, 2)] = "9"
keypad[(1, 1)] = "A"
keypad[(2, 1)] = "B"
keypad[(3, 1)] = "C"
keypad[(2, 0)] = "D"

x, y = 0, 2
for digit in list:
   for i in range(0, len(digit)):
      if digit[i] == "U":
         if (x, y+1) in keypad:
            y += 1
      elif digit[i] == "R":
         if(x+1, y) in keypad:
            x += 1
      elif digit[i] == "D":
         if(x, y-1) in keypad:
            y -= 1
      else:
         if(x-1, y) in keypad:
            x -= 1
   print(keypad[(x, y)])