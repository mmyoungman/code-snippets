def callback(val):
  print("Function callback was called with {0}".format(val))

def caller(val, func):
  func(val)

for i in range(5):
  caller(i, callback)