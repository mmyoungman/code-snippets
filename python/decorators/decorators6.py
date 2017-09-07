def amendFunc(func):
  def newAddFunc(x,y):
    print("Are we going to add?")
    result = func(x,y) # call add
    print("We've done it!")
    return result
  return newFunc

@amendFunc
def add(x,y):
  return x + y

print(add(2,1))