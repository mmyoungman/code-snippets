import time

t0 = time.clock()

total = 0
for i in range(1000000000):
  total += i

t1 = time.clock()

totalTime = t1-t0

print(totalTime)