steps = 316
buffer = [0]
pos = 0

for i in range(1, 2018):
   pos = (pos + steps) % len(buffer)
   pos += 1
   buffer.insert(pos, i)

for i in range(len(buffer)):
   if buffer[i] == 2017:
      print(buffer[(i+1) % len(buffer)])