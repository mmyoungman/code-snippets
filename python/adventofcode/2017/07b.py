import re

file = open("07input.txt", 'r')

list = []

while True:
  line = file.readline()
  if line == '':
    break
  line = line.rstrip('\n')
  list.append(line)

file.close()

def getValue(node, populatedTree):
   if len(node) < 3:
      return int(node[1])
   else:
      result = int(node[1])
      for i in range(2, len(node)):
         for j in range(0, len(populatedTree)):
            if node[i] == populatedTree[j][0]:
               result += getValue(populatedTree[j], populatedTree)
               break
      return result

tree = []

for item in list:
   split = item.split()
   node = []
   node.append(split[0])
   for i in range(1, len(split)):
      split[i] = split[i].strip(',()')
      if split[i].isalnum():
         node.append(split[i])
   tree.append(node)

for i in range(0, len(tree)):
   if len(tree[i]) < 4:   # Only interested if there are two or more branches
      continue
   else:
      values = []
      for j in range(2, len(tree[i])):
         for k in range(0, len(tree)):
            if tree[i][j] == tree[k][0]:
               values.append(getValue(tree[k], tree))
               break
      for j in range(1, len(values)):
         if values[0] != values[j]:
            print(tree[i], values)
            break

# From the output you can deduce that lnpuarm has the wrong weight, too heavy by 8