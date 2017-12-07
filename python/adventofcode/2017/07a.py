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

# print(list)

tree = []

for item in list:
   split = item.split()
   node = []
   node.append(split[0])
   for i in range(1, len(split)):
      split[i] = split[i].strip(',')
      if split[i].isalpha():
         node.append(split[i])
   tree.append(node)

index = 0
newIndex = -1

while index != newIndex:
   index = newIndex
   for i in range(0, len(tree)):
      if tree[index][0] in tree[i][1:]:
         newIndex = i
         break
  
print(tree[index][0])