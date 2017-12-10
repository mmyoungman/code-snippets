file = open("08input.txt", 'r')

list = []

while True:
  line = file.readline()
  if line == '':
    break
  line = line.rstrip('\n')
  list.append(line.split())

file.close()

registers = {}
maxReg = 0

for line in list:
   if line[0] not in registers:
      registers[line[0]] = 0
   
   if line[4] not in registers:
      registers[line[4]] = 0

   if eval(str(registers[line[4]]) + ' ' + line[5] + ' ' + line[6]):
      if line[1] == "inc":
         registers[line[0]] += int(line[2])
      else:
         registers[line[0]] -= int(line[2])
   
   if registers[max(registers, key=registers.get)] > maxReg:
      maxReg = registers[max(registers, key=registers.get)]

print(maxReg)