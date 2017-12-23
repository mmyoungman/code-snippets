from math import sqrt
from itertools import count, islice

def is_prime(n):
   return n > 1 and all(n % i for i in islice(count(2), int(sqrt(n)-1)))

b = 106500
c = 123500
h = 0

# Doesn't work -- results in 916 rather than 917 because 123500 isn't tested!
# while b != c:
#    if not is_prime(b):
#       h += 1
#    b += 17

for n in range(106500, 123501, 17):
   if not is_prime(n):
      h += 1

print(h)