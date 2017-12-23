file = open("21input.txt", 'r')

rules = []

while True:
  line = file.readline()
  if line == '':
    break
  line = line.rstrip('\n')
  rules.append(line.split(' '))

file.close()

for i in range(len(rules)):
   rules[i][0] = rules[i][0].split('/')
   del rules[i][1]
   rules[i][1] = rules[i][1].split('/')

grid = ['.#.', '..#', '###']

def flip(pattern):
   newpat = pattern[:]
   for i in range(len(pattern)):
      row = ''
      for j in range(len(pattern[i])):
         row += pattern[i][(len(pattern[i])-1) - j]
      newpat[i] = row
   return newpat

def rotate(pattern):
   newpat = pattern[:]
   for j in range(len(pattern)):
      row = ''
      for i in range(len(pattern)):
         row += pattern[i][(len(pattern)-1) - j]
      newpat[j] = row
   return newpat

def match(pattern, square):
   if len(pattern) != len(square):
      return False
   for _ in range(4):
      if pattern == square:
         return True
      pattern = rotate(pattern)
   pattern = flip(pattern)
   for _ in range(3):
      if pattern == square:
         return True
      pattern = rotate(pattern)
   if pattern == square:
      return True
   return False

iterations = 5
pix = 0
for _ in range(iterations):
   if len(grid) % 2 == 0:
      pix = 2
   else:
      pix = 3
   newgrid = []
   for y in range(len(grid)//pix):
      for x in range(len(grid)//pix):
         subgrid = grid[y*pix:(y*pix)+pix]
         for i in range(pix):
            subgrid[i] = subgrid[i][x*pix:(x*pix)+pix]
         for rule in rules:
            if match(rule[0], subgrid):
               if x == 0:
                  newgrid += rule[1]
               else:
                  for i in range(pix+1):
                     newgrid[(y*(pix+1)) + i] += rule[1][i]
               break

   grid = newgrid

pixels = 0
for i in range(len(grid)):
   for j in range(len(grid)):
      if grid[i][j] == '#':
         pixels += 1

print(pixels)