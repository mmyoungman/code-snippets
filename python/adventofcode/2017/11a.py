file = open("11input.txt", 'r')

list = []

while True:
  line = file.readline()
  if line == '':
    break
  line = line.rstrip('\n')
  list = line.split(',')

file.close()

x, y, z = 0, 0, 0

n = (0, 1, -1)
ne = (1, 0, -1)
se = (1, -1, 0)
s = (0, -1, 1)
sw = (-1, 0, 1)
nw = (-1, 1, 0)

for dir in list:
   if dir == "n":
      y += 1
      z -= 1
   if dir == "ne":
      x += 1
      z -= 1
   if dir == "se":
      x += 1
      y -= 1
   if dir == "s":
      y -= 1
      z += 1
   if dir == "sw":
      x -= 1
      z += 1
   if dir == "nw":
      x -= 1
      y += 1

print( (abs(x) + abs(y) + abs(z)) / 2 )