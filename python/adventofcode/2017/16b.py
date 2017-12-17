import sys
file = open("16input.txt", 'r')

list = []

list = file.readline().rstrip('\n').split(',')

file.close()

programs = []
char = ord('a')
for i in range(16): 
   programs.append(chr(char + i))

prevConfig = ["abcdefghijklmnop"]
runs = 1000000000

for run in range(runs):
   for command in list:
      if command[0] == 's':
         spin = int(command[1:])
         programs = programs[-spin:] + programs[:-spin]

      if command[0] == 'x':
         toSwap = command[1:].split('/')
         programs[int(toSwap[0])], programs[int(toSwap[1])] = programs[int(toSwap[1])], programs[int(toSwap[0])]

      if command[0] == 'p':
         toSwap = command[1:].split('/')
         for i in range(len(programs)):
            if programs[i] == toSwap[0]:
               toSwap[0] = i
            if programs[i] == toSwap[1]:
               toSwap[1] = i
         programs[toSwap[0]], programs[toSwap[1]] = programs[toSwap[1]], programs[toSwap[0]]

   if ''.join(programs) == "abcdefghijklmnop":
      print(prevConfig[runs % len(prevConfig)])
      sys.exit(0)
   else:
      prevConfig.append(''.join(programs))