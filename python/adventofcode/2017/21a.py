import sys

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

print(rules)

grid = ['.#.', '..#', '###']

def flip(pattern):
   newpat = pattern
   for i in range(len(pattern)):
      for j in range(len(pattern[i])):
         newpat[i][j] = pattern[i][len(pattern[i]) - j]

def rotate(pattern):
   newpat = pattern
   for i in range(len(pattern)):
      for j in range(len(pattern)):
         newpat[j][i] = pattern[i][(len(pattern)-1) - j]


def match(pattern, square):
   if len(pattern) != len(square):
      return False
   if pattern == square:
      return True
   pattern = rotate(pattern)
   if pattern == square:
      return True
   pattern = rotate(pattern)
   if pattern == square:
      return True
   pattern = rotate(pattern)
   if pattern == square:
      return True
   pattern = rotate(pattern)
   pattern = flip(pattern)
   if pattern == square:
      return True
   pattern = rotate(pattern)
   if pattern == square:
      return True
   pattern = rotate(pattern)
   if pattern == square:
      return True
   pattern = rotate(pattern)
   if pattern == square:
      return True

   return False

iterations = 5
for _ in range(iterations):
   if len(grid) % 2 == 0:
      pix = 2
   else:
      pix = 3
   newgrid = []
   for y in range(len(grid)//pix):
      for x in range(len(grid)//pix):
         # matchFound = False
         for rule in rules:
            if match(rule[0], grid[int(y*pix):int((y*pix)+pix)][int(x*pix):int((x*pix)+pix)]):
               # matchFound = True
               if x == 0:
                  newgrid.append(rule[1]) 
               else:
                  for i in range(pix+1):
                     newgrid[(y*(pix+1)) + i] += rule[1][i]
               break
         # if not matchFound:
         #    print("Didn't find a match!")

   grid = newgrid

pixels = 0
for i in range(len(grid)):
   for j in range(len(grid)):
      if grid[i][j] == '#':
         pixels += 1

print(pixels)