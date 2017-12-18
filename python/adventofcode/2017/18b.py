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

def getvalue(arg, values):
   try:
      value = int(arg)
   except:
      if arg not in values.keys():
         values[arg] = 0
      value = values[arg] 
   return value

def tick(pos, values, sndmsgs, rcvmsgs):
   if cmds[pos][0] == "set":
      values[cmds[pos][1]] = getvalue(cmds[pos][2], values)
   elif cmds[pos][0] == "add":
      values[cmds[pos][1]] += getvalue(cmds[pos][2], values)
   elif cmds[pos][0] == "mul":
      values[cmds[pos][1]] *= getvalue(cmds[pos][2], values)
   elif cmds[pos][0] == "mod":
      values[cmds[pos][1]] %= getvalue(cmds[pos][2], values)
   elif cmds[pos][0] == "jgz":
      if getvalue(cmds[pos][1], values) > 0:
         pos += getvalue(cmds[pos][2], values) - 1
   elif cmds[pos][0] == "snd":
      sndmsgs.append(getvalue(cmds[pos][1], values))
   elif cmds[pos][0] == "rcv":
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
      print(sentMsgs0)
      sys.exit(0)