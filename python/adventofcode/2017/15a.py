genA = 703
genB = 516

count = 0

for i in range(40000000):
   genA = (genA * 16807) % 2147483647
   genB = (genB * 48271) % 2147483647

   binA = format(genA, '016b')
   binB = format(genB, '016b')

   # print(genA)
   # print(binA)
   # print(binA[len(binA)-16:len(binA)])

   if binA[len(binA)-16:len(binA)] == binB[len(binB)-16:len(binB)]:
      count += 1
   
print(count)
