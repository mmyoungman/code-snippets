file = open("22input.txt", 'r')

map = []

while True:
  line = file.readline()
  if line == '':
    break
  line = line.rstrip('\n')
  map.append([char for char in line])

file.close()

posX = len(map)//2
posY = len(map)//2
dir = 0 # 0 is north, 1 east, 2 south, 3 west
count = 0

for _ in range(10000000):
   if map[posY][posX] == '#':
      map[posY][posX] = 'F'
      dir = (dir+1) % 4
   elif map[posY][posX] == '.':
      map[posY][posX] = 'W'
      dir = (dir-1) % 4
   elif map[posY][posX] == 'W':
      count += 1
      map[posY][posX] = '#'
   else:
      map[posY][posX] = '.'
      dir = (dir+2) % 4

   if dir == 0:
      posY -= 1
   elif dir == 1:
      posX += 1
   elif dir == 2:
      posY += 1
   else:
      posX -= 1

   if posX == 0:
      for row in range(len(map)):
         map[row].insert(0, '.') 
         posX = 1
   elif posX == len(map)-1:
      for row in range(len(map)):
         map[row].append('.')
   
   if posY == 0:
      map.insert(0, ['.' for _ in range(len(map[0]))])
      posY = 1
   if posY == len(map)-1:
      map.append(['.' for _ in range(len(map[0]))])

print(count)