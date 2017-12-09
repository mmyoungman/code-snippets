file = open("09input.txt", 'r')

list = []

while True:
  line = file.readline()
  if line == '':
    break
  line = line.rstrip('\n')
  list.append(line)

file.close()

rubbishCount = 0

for line in list:
   skipNext = False
   rubbish = False
   for char in line:
      if skipNext:
         skipNext = False
         continue
      if char == "!":
         skipNext = True
         continue
      if rubbish:
         if char == ">":
            rubbish = False
            continue
         rubbishCount += 1
         continue
      if char == "<":
         rubbish = True
         continue

print(rubbishCount)