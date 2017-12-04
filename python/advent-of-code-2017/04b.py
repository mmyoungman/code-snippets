import re

file = open("04input.txt", 'r')

list = []

while True:
  line = file.readline()
  if line == '':
    break
  list.append(re.split(r' ', line.rstrip('\n')))

file.close()

numValid = 0

for row in list:
   phrases = []
   isValid = True
   for phrase in row:
      phrase = sorted(phrase)
      if phrase in phrases:
         isValid = False
         break
      else:
         phrases.append(phrase)
   if isValid:
      numValid += 1

print(numValid)