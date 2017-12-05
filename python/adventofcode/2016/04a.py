import re

file = open("04input.txt", 'r')

list = []

while True:
  line = file.readline()
  if line == '':
    break
  line = line.rstrip('\n')
  list.append(line)

file.close()

result = 0

for room in list:
   checksum = re.search(r"\[[a-z]+\]", room).group(0).strip('[]')
   sector_id = re.search(r"[0-9]+", room).group(0)
   room = room.split('-')
   del room[len(room)-1]
   room = ''.join(room)

   dict = {}
   for char in room:
      if char not in dict:
         dict[char] = 1
      else:
         dict[char] += 1
   roomKey = sorted(dict, key=dict.__getitem__, reverse=True)
   for i in range(len(roomKey)):
      sameRange = i+1
      while sameRange < len(roomKey) and dict[roomKey[i]] == dict[roomKey[sameRange]]:
         sameRange += 1
      innerList = sorted(roomKey[i:sameRange])

      index = 0
      for j in range(i, sameRange):
         roomKey[j] = innerList[index]
         index += 1

   if ''.join(roomKey[:5]) == checksum:
      result += int(sector_id)

print(result)