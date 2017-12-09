file = open("09input.txt", 'r')

list = []

while True:
  line = file.readline()
  if line == '':
    break
  line = line.rstrip('\n')
  list.append(line)

file.close()

total = 0

for line in list:
   depth = 0
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
      if char == "<":
         rubbish = True
         continue
      if char == "{":
         depth += 1
         total += depth
         continue
      if char == "}":
         depth -= 1

print(total)