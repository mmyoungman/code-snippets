genA = 703
genB = 516

count = 0
judgeA = []
judgeB = []

while True:
   genA = (genA * 16807) % 2147483647
   genB = (genB * 48271) % 2147483647

   if genA % 4 == 0:
      judgeA.append(genA)

   if genB % 8 == 0:
      judgeB.append(genB)
   
   if len(judgeA) >= 5000000 and len(judgeB) >= 5000000:
      break

print(len(judgeA), len(judgeB))

if len(judgeA) < len(judgeB):
   for i in range(len(judgeA)):
      binA = format(judgeA[i], '016b')
      binB = format(judgeB[i], '016b')
      if binA[len(binA)-16:len(binA)] == binB[len(binB)-16:len(binB)]:
         count += 1
else:
   for i in range(len(judgeB)):
      binA = format(judgeA[i], '016b')
      binB = format(judgeB[i], '016b')
      if binA[len(binA)-16:len(binA)] == binB[len(binB)-16:len(binB)]:
         count += 1
   
print(count)
