class repeater:
  def __init__(self, value):
    self.value = value

  def __iter__(self):
    return repeaterIterator(self)

class repeaterIterator:
  def __init__(self, source):
    self.source = source

  def __next__(self):
    return self.source.value

rep = repeater('hello')
for i in rep:
  print(i)

"""
rep2 = repeater('hello')
iterator = rep2.__iter__()
while True:
  i = iterator.__next__()
  print(i)
"""