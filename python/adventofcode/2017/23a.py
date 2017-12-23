file = open("23inputb.txt", 'r')

cmds = []

while True:
   line = file.readline()
   if line == '':
      break
   cmds.append(line.rstrip('\n').split())

file.close()

def getvalue(arg, values):
   try:
      value = int(arg)
   except:
      # if arg not in values.keys():
      #    values[arg] = 0
      value = values[arg] 
   return value

registers = {}
for char in "abcdefgh":
   registers[char] = 0

pos = 0
count = 0

while True:
   if cmds[pos][0] == "set":
      registers[cmds[pos][1]] = getvalue(cmds[pos][2], registers)
   elif cmds[pos][0] == "sub":
      registers[cmds[pos][1]] -= getvalue(cmds[pos][2], registers)
   elif cmds[pos][0] == "mul":
      count += 1
      registers[cmds[pos][1]] *= getvalue(cmds[pos][2], registers)
   elif cmds[pos][0] == "jnz":
      if getvalue(cmds[pos][1], registers) != 0:
         pos += getvalue(cmds[pos][2], registers) - 1
   pos += 1

   if pos < 0 or pos >= len(cmds):
      break

print(count)