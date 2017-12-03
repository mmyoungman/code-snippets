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
    value = int(list[i][j])
    if j == 0:
      low = value
      high = value
    else:
      if value < low:
        low = value
      if value > high:
        high = value
  result += high - low

print(result)
