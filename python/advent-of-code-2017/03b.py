import sys

puzzleinput = 312051

x, y = 0, 0
incLength = False
count = 1
length = 1
direction = "S"
spiral = {}

spiral[(x, y)] = count
while True:
   if direction == "S":
      direction = "E"
   elif direction == "E":
      direction = "N"
   elif direction == "N":
      direction = "W"
   else:
      direction = "S"
   
   target = count + length
   while count < target: 
      if direction == "N":
         y += 1
      elif direction == "E":
         x += 1
      elif direction == "S":
         y -= 1
      else:
         x -= 1
      count += 1

      value = 0
      if (x+1, y) in spiral:
         value += spiral[(x+1, y)]
      if (x+1, y+1) in spiral:
         value += spiral[(x+1, y+1)]
      if (x, y+1) in spiral:
         value += spiral[(x, y+1)]
      if (x-1, y+1) in spiral:
         value += spiral[(x-1, y+1)]
      if (x-1, y) in spiral:
         value += spiral[(x-1, y)]
      if (x-1, y-1) in spiral:
         value += spiral[(x-1, y-1)]
      if (x, y-1) in spiral:
         value += spiral[(x, y-1)]
      if (x+1, y-1) in spiral:
         value += spiral[(x+1, y-1)]
         
      spiral[(x, y)] = value

      if value > puzzleinput:
         print(value)
         sys.exit(0)

   if incLength == True:
      length += 1
   incLength = not incLength