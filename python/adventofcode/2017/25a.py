tape = [0]
state = 'A'
index = 0

for _ in range(12399302):
   if state == 'A':
      if tape[index] == 0:
         tape[index] = 1
         index += 1
         state = 'B'
      else:
         tape[index] = 0
         index += 1
         state = 'C'
   elif state == 'B':
      if tape[index] == 0:
         index -= 1
         state = 'A'
      else:
         tape[index] = 0
         index += 1
         state = 'D'
   elif state == 'C':
      if tape[index] == 0:
         tape[index] = 1
         index += 1
         state = 'D'
      else:
         index += 1
         state = 'A'
   elif state == 'D':
      if tape[index] == 0:
         tape[index] = 1
         index -= 1
         state = 'E'
      else:
         tape[index] = 0
         index -= 1
         state = 'D'
   elif state == 'E':
      if tape[index] == 0:
         tape[index] = 1
         index += 1
         state = 'F'
      else:
         index -= 1
         state = 'B'
   elif state == 'F':
      if tape[index] == 0:
         tape[index] = 1
         index += 1
         state = 'A'
      else:
         index += 1
         state = 'E'
   
   if index == -1:
      tape.insert(0, 0)
      index = 0
   
   if index == len(tape):
      tape.append(0)

result = 0
for bit in tape:
   if bit == 1:
      result += 1
print(result)