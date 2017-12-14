input = "ugkiagan"
grid = []

for i in range(128):
   rowInput = input + '-' + str(i)
   rowList = []
   for char in rowInput:
      rowList.append(ord(char))
   rowList += [17,31,73,47,23]

   numList = []
   for j in range(256):
      numList.append(j)

   pos = 0
   skip = 0
   for index in range(0, 64):
      for jumpLen in rowList:
         subList = []
         if jumpLen < 2:
            pass
         elif (pos+jumpLen)%256 > pos%256:
            subList = numList[pos%256:(pos+jumpLen)%256]
            subList = subList[::-1]
            j = 0
            for i in range(pos%256, (pos+jumpLen)%256):
               numList[i] = subList[j]
               j += 1
         else:
            subList = numList[pos%256:256] + numList[0:(pos+jumpLen)%256]
            subList = subList[::-1]
            j = 0
            for i in range(pos%256, 256):
               numList[i] = subList[j]
               j += 1
            for i in range(0, (pos+jumpLen)%256):
               numList[i] = subList[j]
               j += 1
         pos += jumpLen + skip
         skip += 1
   result = []
   for i in range(0, 256, 16):
      newValue = 0
      for j in range(i, i+16):
         newValue ^= numList[j]
      result.append('{:08b}'.format(newValue))

   grid.append(''.join(result))

gridDict = {}
for y in range(128):
   for x in range(128):
      if grid[y][x] == "1":
         gridDict[(x, y)] = grid[y][x]

def zeroRegion(x, y):
   del gridDict[(x, y)]
   if (x-1, y) in gridDict:
      zeroRegion(x-1, y)
   if (x+1, y) in gridDict:
      zeroRegion(x+1, y)
   if (x, y-1) in gridDict:
      zeroRegion(x, y-1)
   if (x, y+1) in gridDict:
      zeroRegion(x, y+1)

regions = 0
for x in range(128):
   for y in range(128):
      if (x, y) in gridDict:
         regions += 1
         zeroRegion(x, y)

print(regions)
