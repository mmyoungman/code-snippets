import sys

puzzleinput = 312051

x, y = 0, 0
incLength = False
count = 1
length = 1
direction = "S"

while count != puzzleinput:
   if direction == "S":
      direction = "E"
   elif direction == "E":
      direction = "N"
   elif direction == "N":
      direction = "W"
   else:
      direction = "S"
   
   if puzzleinput < count + length:
      while count < puzzleinput:
         if direction == "N":
            y += 1
         elif direction == "E":
            x += 1
         elif direction == "S":
            y -= 1
         else:
            x -= 1
         count += 1
      break

   count += length
   if direction == "N":
      y += length
   elif direction == "E":
      x += length
   elif direction == "S":
      y -= length
   else:
      x -= length
   
   if incLength == True:
      length += 1
   incLength = not incLength

print(abs(x) + abs(y))