file = open("10input.txt", 'r')

list = []

while True:
  line = file.readline()
  if line == '':
    break
  line = line.rstrip('\n')
  list = line.split(',')

file.close()

for i in range(0, len(list)):
   list[i] = int(list[i])

numList = []
for i in range(0, 256):
   numList.append(i)

pos = 0
skip = 0

for jumpLen in list:
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

print(numList[0] * numList[1])