# NOTE: Doesn't work!

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

values0 = {'p': 0}
values1 = {'p': 1}

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
         values[cmds[pos]     if len(rcvmsgs) < 1:   
[1]] = 0
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
      sndmsgs.append(values[cmds[pos][1]])

   if cmds[pos][0] == "rcv":
      if len(rcvmsgs) < 1:   
         return pos, values, sndmsgs, rcvmsgs, False
      else:
         values[cmds[pos][1]] = rcvmsgs.pop(0)

   pos += 1
   
   return pos, values, sndmsgs, rcvmsgs, True

sentMsgs0 = 0

while True:
   pos0, values0, msgs1, msgs0, changed0 = tick(pos0, values0, msgs1, msgs0)
   old = len(msgs0)
   pos1, values1, msgs0, msgs1, changed1 = tick(pos1, values1, msgs0, msgs1)
   if len(msgs0) > old:
      sentMsgs0 += 1 
   if (not changed0) and (not changed1):
      print("DEADLOCK!")
      print(pos0, pos1)
      print(sentMsgs0)
      sys.exit(0)