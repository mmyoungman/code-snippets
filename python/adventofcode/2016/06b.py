import re

file = open("06input.txt", 'r')

list = []

while True:
  line = file.readline()
  if line == '':
    break
  line = line.rstrip('\n')
  list.append(line)

file.close()

print(list)
msg = ""
for i in range( len(list[0]) ):
   dict = {}
   for word in list:
      if word[i] in dict:
         dict[word[i]] += 1
      else:
         dict[word[i]] = 1
   msg += min(dict, key=dict.get)

print(msg)