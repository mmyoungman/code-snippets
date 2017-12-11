file = open("10input.txt", 'r')

list = []

while True:
  line = file.readline()
  if line == '':
    break
  line = line.rstrip('\n')
  for char in line:
     list.append(ord(char))

file.close()

list = list + [17, 31, 73, 47, 23]

numList = []
for i in range(0, 256):
   numList.append(i)

pos = 0
skip = 0

for index in range(0, 64):
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

result = []

for i in range(0, 256, 16):
   newValue = 0
   for j in range(i, i+16):
      newValue ^= numList[j]
   result.append('{:02x}'.format(newValue))

print(''.join(result))