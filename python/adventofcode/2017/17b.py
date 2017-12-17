steps = 316
pos = 0
result = 1

for i in range(1, 50000000):
   pos = (pos + steps) % i
   pos += 1
   if pos == 1:
      result = i

print(result)