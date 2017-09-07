class boundedRepeater:
  def __init__(self, value, numRepeats):
    self.value = value
    self.numRepeats = numRepeats
    self.count = 0

  def __iter__(self):
    return self

  def __next__(self):
    if self.count >= self.numRepeats
      raise StopIteration
    self.count += 1
    return self.value

rep = boundedRepeater('hello', 4)
for i in rep:
  print(i)

"""
rep2 = boundedRepeater('hello', 4)
iterator = iter(rep2)
while True:
  try:
    i = next(iterator)
  except StopIteration:
    break
  print(i)
"""