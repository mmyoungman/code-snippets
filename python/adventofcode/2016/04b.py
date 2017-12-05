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

for room in list:
   checksum = re.search(r"\[[a-z]+\]", room).group(0).strip('[]')
   sector_id = re.search(r"[0-9]+", room).group(0)
   room = room.split('-')
   del room[len(room)-1]

   shift = int(sector_id) % 26

   for i in range(0, len(room)):
      newWord = ""
      for j in range(0, len(room[i])):
         value = (ord(room[i][j]) - 97) + shift
         if value > 25:
            value -= 26
         newWord += chr(97 + value)
      room[i] = newWord

   room = ' '.join(room)
   print(room, sector_id)