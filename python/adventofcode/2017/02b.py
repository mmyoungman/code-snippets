import re

file = open("02input.txt", 'r')

list = []

while True:
  line = file.readline()
  if line == '':
    break
  list.append(re.split(r'\t+', line.rstrip('\n')))

file.close()

result = 0

for i in range(0, len(list)):
  for j in range(0, len(list[i])):
    for k in range(0, len(list[i])):
      if j == k:
        continue
      value1 = int(list[i][j])
      value2 = int(list[i][k])
      if value1 % value2 == 0:
        result += value1 // value2
        break
        
print(result)