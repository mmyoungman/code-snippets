import sys

file = open("19input.txt", 'r')

list = []

while True:
  line = file.readline()
  if line == '':
    break
  line = line.rstrip('\n')
  list.append(line)

file.close()

posx = 0
posy = 0

for x in range(len(list[0])):
   if list[0][x] == "|":
      posx = x

dir = "S"
steps = 1  # Because we've already covered the first step

while True:
   if dir == 'N':
      posy -= 1
   elif dir == 'S':
      posy += 1
   elif dir == 'E':
      posx += 1
   else:
      posx -= 1

   if list[posy][posx] == ' ':
      print(steps)
      sys.exit(0)
   elif list[posy][posx] == '+':
      if dir == 'N' or dir == 'S':
         if list[posy][posx+1] != ' ':
            dir = 'E'
         else:
            dir = 'W'
      elif dir == 'E' or dir == 'W':
         if list[posy+1][posx] != ' ':
            dir = 'S'
         else:
            dir = 'N'
   steps += 1