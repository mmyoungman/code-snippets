import sys

puzzleinput = 312051

x, y = 0, 0
incLength = False
count = 1
length = 1
direction = "S"
spiral = {}

spiral[(x, y)] = count
while count != puzzleinput:
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
      spiral[(x, y)] = count
      if count == puzzleinput:
         print(abs(x) + abs(y))
         sys.exit(0)

   if incLength == True:
      length += 1
   incLength = not incLength