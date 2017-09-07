# This program adds up integers given to the program as command line arguments

import sys

try:
  total = sum(int(arg) for arg in sys.argv[1:])
  print('sum = ', total)
except ValueError:
  print('Please supply integer arguments')