import sys
file = open("18input.txt", 'r')

cmds = []

while True:
   line = file.readline()
   if line == '':
      break
   cmds.append(line.rstrip('\n').split())

file.close()

pos0 = 0
pos1 = 0

values0 = {}
values1 = {}

msgs0 = []
msgs1 = []

sound = 0

def tick(pos, values, sndmsgs, rcvmsgs):
   global sound
   if cmds[pos][0] == "set":
      if cmds[pos][1] not in values.keys():
         values[cmds[pos][1]] = 0
      if cmds[pos][2].lstrip('-').isdigit():
         values[cmds[pos][1]] = int(cmds[pos][2])
      else:
         if cmds[pos][2] not in values.keys():
            values[cmds[pos][2]] = 0
         values[cmds[pos][1]] = values[cmds[pos][2]]
   
   if cmds[pos][0] == "add":
      if cmds[pos][1] not in values.keys():
         values[cmds[pos][1]] = 0
      if cmds[pos][2].lstrip('-').isdigit():
         values[cmds[pos][1]] += int(cmds[pos][2])
      else:
         if cmds[pos][2] not in values.keys():
            values[cmds[pos][2]] = 0
         values[cmds[pos][1]] += values[cmds[pos][2]]

   if cmds[pos][0] == "mul":
      if cmds[pos][1] not in values.keys():
         values[cmds[pos][1]] = 0
      if cmds[pos][2].lstrip('-').isdigit():
         values[cmds[pos][1]] *= int(cmds[pos][2])
      else:
         if cmds[pos][2] not in values.keys():
            values[cmds[pos][2]] = 0
         values[cmds[pos][1]] *= values[cmds[pos][2]]

   if cmds[pos][0] == "mod":
      if cmds[pos][1] not in values.keys():
         values[cmds[pos][1]] = 0
      if cmds[pos][2].lstrip('-').isdigit():
         values[cmds[pos][1]] %= int(cmds[pos][2])
      else:
         if cmds[pos][2] not in values.keys():
            values[cmds[pos][2]] = 0
         values[cmds[pos][1]] %= values[cmds[pos][2]]
   
   if cmds[pos][0] == "jgz":
      jmp = False
      if cmds[pos][1].lstrip('-').isdigit():
         if int(cmds[pos][1]) > 0:
            jmp = True
      elif values[cmds[pos][1]] > 0:
         jmp = True
      
      if jmp:
         if cmds[pos][2].lstrip('-').isdigit():
            pos += int(cmds[pos][2]) - 1
         else:
            if cmds[pos][2] not in values.keys():
               values[cmds[pos][2]] = 0
            pos += values[cmds[pos][2]] - 1

   if cmds[pos][0] == "snd":
      if cmds[pos][1] not in values.keys():
         values[cmds[pos][1]] = 0
      sound = values[cmds[pos][1]]

   if cmds[pos][0] == "rcv":
      if sound != 0:
         print(sound)
         sys.exit(0)
   
   pos += 1
   
   return pos, values, msgs1, msgs0

while True:
   pos0, values0, msgs1, msgs0 = tick(pos0, values0, msgs1, msgs0)