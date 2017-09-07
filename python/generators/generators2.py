class repeater:
  def __init__(self, value):
    self.value = value

  def __iter__(self):
    return self

  def __next__(self):
    return self.value

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